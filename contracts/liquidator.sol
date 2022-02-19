// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./interface.sol";
import "./PancakeLibrary.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract UniFlashSwap is IPancakeCallee,Ownable{
    address private constant ComptrollerAddr = 0xfD36E2c2a6789Db23113685031d7F16329158384;
    address private constant wBNB = 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c;
    address private constant FACTORY = 0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73;
    address private constant vBNB = 0xA07c5b74C9B40447a954e1466938b865b6BBea36;
    address private constant vBUSD = 0x95c78222B3D6e262426483D42CfA53685A67Ab9D;
    address private constant vUSDT = 0xfD5840Cd36d94D7229439859C0112a4185BC0255;
    address private constant vDAI = 0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1;
    address private constant ROUTER = 0x10ED43C718714eb63d5aA57B78B54704E256024E;
    address private constant VAI = 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7;
    address private constant USDT = 0x55d398326f99059fF775485246999027B3197955;
    uint private constant MAXUINT32 = ~uint(0);

    mapping(address => mapping(address => bool)) approves;

    event Scenario(uint scenarioNo, address repayUnderlyingToken, uint repayAmount, address seizedUnderlyingToken, uint flashLoanReturnAmount,uint seizedUnderlyingAmount, uint massProfit);
    event SeizedVTokenAmount(uint, uint, uint);
    event SeizedUnderlyingTokenAmount(uint, uint);
    event NotEnoughSeizedUnderlygingAmount(uint, uint);

    event Withdraw(address indexed, address indexed, uint);
    event WithdrawETH(address indexed, uint);
    event Qingsuan(uint, uint);

    struct LocalVars {
        uint situation;
        address flashLoanFrom;
        address[] path1;
        address[] path2;
        address[] tokens;
        uint repayAmount;
        uint flashLoanReturnAmount;
        address borrower;

        //vToken
        uint seizedVTokenAmount;

        //underlyingToken
        uint seizedUnderlyingAmount;
        
        uint massProfit;
    }

    function swapOneBNBToFlashLoandUnderlyingToken(address _flashLoanUnderlyingToken) onlyOwner public{
        uint amount = 1 ether;
         IWETH(wBNB).deposit{value: amount}();

        address[] memory path = new address[](2);
        path[0] = wBNB;
        path[1] = _flashLoanUnderlyingToken; 
        chainSwapExactIn(amount, path, address(this));
    }

    //situcation： 情况 1-5
    //ch： 借钱用的pair地址
    //sellPath： 卖的时候的path
    //tokens：
    // Tokens array
    // [0] - _flashLoanVToken 要去借的钱（要还给venus的）
    // [1] - _seizedVToken 可以赎回来的钱
    // [2] - _seizedTokenUnderlying 赎回来的钱的underlying
    // [3] - _flashloanTokenUnderlying 借的钱的underlying
    // [4] - target 目标账号
    //_flashLoanAmount ： 借多少？ 还多少？
    /*
    case1 
       0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE
       ["0xfD5840Cd36d94D7229439859C0112a4185BC0255","0xfD5840Cd36d94D7229439859C0112a4185BC0255","0x55d398326f99059fF775485246999027B3197955","0x55d398326f99059fF775485246999027B3197955","0x564EE8bF0bA977A1ccc92fe3D683AbF4569c9f5E"]
       21034205123917372652

    case2.2 
        height15379782, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, repaySmbol:vETH, flashLoanFrom:0x74E4716E431f45807DCF19f284c7aA99F18a4fbc, repayAddress:0xf508fCD89b8bd15579dc79A6827cB4686A3592c8, repayValue:27760775212595059202.7, repayAmount:9834830202498401 seizedSymbol:vETH, seizedAddress:0xf508fCD89b8bd15579dc79A6827cB4686A3592c8, seizedCTokenAmount:53542165, seizedUnderlyingTokenAmount:10818313118937201.583675475128906, seizedUnderlyingTokenValue:30536852440824038910.2407636463629662
        calculateSeizedTokenAmount case2: seizedSymbol == repaySymbol and symbol is not stable coin, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, symbol:vETH, seizedAmount:10818313118937201.583675475128906, returnAmout:9859478899747771, usdtAmount:2702645379426654958, gasFee:2425666800000000000, profit:0.2786001639516656
        case2, profitable liquidation catched:&{0xFAbE4C180b6eDad32eA0Cf56587c54417189e422 0.9512268623785098 15379543 0001-01-01 00:00:00 +0000 UTC}, profit:0.2786001639516656

        0x74E4716E431f45807DCF19f284c7aA99F18a4fbc
        path1: ["0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x55d398326f99059fF775485246999027B3197955"] 
        path2: ["0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x55d398326f99059fF775485246999027B3197955"]
        tokens:["0xf508fCD89b8bd15579dc79A6827cB4686A3592c8","0xf508fCD89b8bd15579dc79A6827cB4686A3592c8","0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x2170ed0880ac9a755fd29b2688956bd959f933f8","0xFAbE4C180b6eDad32eA0Cf56587c54417189e422"]
        9834830202498401

    */
    function qingsuan(uint _situation, address _flashLoanFrom, address [] calldata  _path1,  address [] calldata  _path2,address [] calldata  _tokens, uint _flashLoanAmount) external {
        require(_situation>=1&&_situation<=7,"wrong si");
        require(_flashLoanFrom != address(0), "!pair");
        {
            (,,uint shortfall) = Comptroller(ComptrollerAddr).getAccountLiquidity(_tokens[4]);
            require(shortfall > 0, "shortfall must greater than zer 0");

            uint borrowBalanceStored = VTokenInterface(_tokens[1]).borrowBalanceStored(_tokens[4]);
            uint closeFactor = Comptroller(ComptrollerAddr).closeFactorMantissa();
            uint maxAllowedRepayAmount = (borrowBalanceStored * closeFactor)/ 1e18;
            require(_flashLoanAmount < maxAllowedRepayAmount, "flashloanAmount must be less than maxAllowedRepayAmount");
        }

        if (!approves[_tokens[3]][_tokens[0]]){
            IERC20(_tokens[3]).approve(_tokens[0], MAXUINT32);
            approves[_tokens[3]][_tokens[0]] = true;
        }

        uint beforeBalance;
        uint afterBalance;

        if (isStableCoin(_tokens[3])){
            beforeBalance = IERC20(_tokens[3]).balanceOf(address(this));
        }else{
            beforeBalance = IERC20(USDT).balanceOf(address(this));
        }

        //token0，token1的顺序要确定好
        address token0 = IPancakePair(_flashLoanFrom).token0();
        address token1 = IPancakePair(_flashLoanFrom).token1();
        //我们只想要一种币，看好0和1那个是我们要借的，把数设置好，另外一种币设置成0
        uint amount0Out = _tokens[3] == token0 ? _flashLoanAmount : 0;
        uint amount1Out = _tokens[3] == token1 ? _flashLoanAmount : 0;
        bytes memory callbackdata = abi.encode(_situation,_flashLoanFrom,_path1,_path2,_tokens,_flashLoanAmount);
        IPancakePair(_flashLoanFrom).swap(amount0Out, amount1Out, address(this), callbackdata);

        if (isStableCoin(_tokens[3])){
            afterBalance = IERC20(_tokens[3]).balanceOf(address(this));
        }else{
            afterBalance = IERC20(USDT).balanceOf(address(this));
        }

        emit Qingsuan(beforeBalance, afterBalance);
    }

    function pancakeCall(
        address _sender,
        uint _amount0,
        uint _amount1,
        bytes calldata _data
    ) external override {
        LocalVars  memory vars;

        (vars.situation,vars.flashLoanFrom, vars.path1,  vars.path2, vars.tokens, vars.repayAmount) = abi.decode(_data, (uint,address,address [],address [],address [],uint));
        require(msg.sender == vars.flashLoanFrom, "!pair");
        require(_sender == address(this), "!sender");

        vars.flashLoanReturnAmount = vars.repayAmount + ((vars.repayAmount * 25) / 9975) + 1;
        vars.borrower = vars.tokens[4];

        IERC20 repayUnderlyingToken = IERC20(vars.tokens[3]);
        IERC20 seizedUnderlyingToken = IERC20(vars.tokens[2]);
        uint[] memory amounts;
        
        if(vars.situation==1){
            //case1: repayToken is USDT, seizedToken is USDT
//            require(vars.path1.length==0 && vars.path2.length==0,"1-patherr");
            require(isStableCoin(vars.tokens[0]), "1-not stable coin");
            require(vars.tokens[0] == vars.tokens[1], "1- not same coin");
            
            (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
            vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);
            require(vars.seizedUnderlyingAmount > vars.flashLoanReturnAmount, "no extra");
            vars.massProfit = vars.seizedUnderlyingAmount - vars.flashLoanReturnAmount;
        }
        else if(vars.situation==2){
            if(isVBNB(vars.tokens[0])) {
                //case2.1 repayToken is BNB, seizedToken is BNB 
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.

                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);
                require(vars.seizedUnderlyingAmount > vars.flashLoanReturnAmount,"2.1-no-extra");

                IWETH(wBNB).deposit{value:vars.seizedUnderlyingAmount}(); //change BNB to wBNB

                uint remain = vars.seizedUnderlyingAmount-vars.flashLoanReturnAmount; //calculate how much wBNB left after return flashloan
                amounts = chainSwapExactIn(remain, vars.path2, address(this));  //swap the left wBNB to USDT
                vars.massProfit = amounts[amounts.length-1];
            }else {
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);
                require(vars.seizedUnderlyingAmount > vars.flashLoanReturnAmount,"2.2-no-extra");
                
                uint remain = vars.seizedUnderlyingAmount - vars.flashLoanReturnAmount;    //calculate how much ETH left after return flashloan
                amounts = chainSwapExactIn(remain,vars.path2,address(this));  //swap the left wETH to USDT
                vars.massProfit = amounts[amounts.length-1];
            }
        }
        else if(vars.situation==3){
            require(isStableCoin(vars.tokens[1]), "3-seized token is not stable coin");
            if (isVBNB(vars.tokens[0])){
                // case3.1 seizedToken is USDT, repayToken is BNB
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.

                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                // change part of USDT to flashLoanReturnAmount wBNB for returning flashloan later
                amounts =  chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "3.1-no-extra");

                vars.massProfit = vars.seizedUnderlyingAmount - amounts[0];
            }else{
                // case3.2 seizedToken is USDT, repayToken is wETH
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                // change part of USDT to flashLoanReturnAmount wETH for returning flashloan later
                amounts = chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "3.2-bnb-no-extra");

                vars.massProfit = vars.seizedUnderlyingAmount - amounts[0];
            }
        }else if(vars.situation==4){
            require(isStableCoin(vars.tokens[0]), "4-repayToken is not stable coin");
            if (isVBNB(vars.tokens[1])){
                //case4.1 seizedToken is BNB, repayToken is USDT
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                IWETH(wBNB).deposit{value:vars.seizedUnderlyingAmount}();  //change BNB to wBNB

                //change all wBNB to USDT
                amounts = chainSwapExactIn(vars.seizedUnderlyingAmount, vars.path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > vars.flashLoanReturnAmount, "4.1-no extra");
                vars.massProfit = usdtAmount - vars.flashLoanReturnAmount;
            }else{
                //case4.2 seizedToken is ETH, repayToken is USDT
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                // change all wETH to USDT
                amounts = chainSwapExactIn(vars.seizedUnderlyingAmount, vars.path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > vars.flashLoanReturnAmount, "4.2-no extra");
                vars.massProfit = usdtAmount - vars.flashLoanReturnAmount;
            }
        }else if(vars.situation==5){
            if (isVBNB(vars.tokens[0])){
                //case5.1 seizedToken is ETH, repayToken is BNB,
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                //change part of wETH to flashLoanReturnAmount wBNB
                amounts = chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change remain wETH to USDT
                uint remain = vars.seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                vars.massProfit = amounts[amounts.length-1];

            }else if (isVBNB(vars.tokens[1])){
                //case5.2 seizedToken is BNB, repayToken is ETH
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                IWETH(wBNB).deposit{value:vars.seizedUnderlyingAmount}();  //change BNB to wBNB

                //change part of wBNB to flashLoanReturnAmount ETH
                amounts = chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change the remained wBNB to USDT
                uint remain = vars.seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                vars.massProfit = amounts[amounts.length-1];
            }else{
                //case5.3 repayToken is wETH, seizedToken is CAKE
                (vars.seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                //change part of CAKE to flashLoanReturnAmount ETH for returning flashloan later
                amounts = chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "5.3-no extra");

                //change the remained CAKE to USDT
                uint remain = vars.seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                vars.massProfit = amounts[amounts.length-1];
            }
        }else if (vars.situation==6){
            require(isStableCoin(vars.tokens[1]), "6-seizedToken is not stable coin");
            //case6 repayToken is VAI, seizedToken is USDT
            uint actualRepayAmount;
            (vars.seizedVTokenAmount, actualRepayAmount) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
            vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

            // change part of USDT to flashLoanReturnAmount VAI for returning flashloan later
            uint changeAmount = vars.flashLoanReturnAmount + actualRepayAmount - vars.repayAmount;
            chainSwapExactOut(changeAmount, vars.path1, address(this));
            require(vars.seizedUnderlyingAmount > amounts[0], "6-noextra");

            vars.massProfit = vars.seizedUnderlyingAmount - amounts[0];
        }else if (vars.situation==7){
            //case7.1 repayToken is VAI, seizedToken is BNB
            if (isVBNB(vars.tokens[1])){
                uint actualRepayAmount;
                (vars.seizedVTokenAmount, actualRepayAmount) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);
                IWETH(wBNB).deposit{value:vars.seizedUnderlyingAmount}();  //change BNB to wBNB

                //change part of wBNB to flashLoanReturnAmount VAI
                uint changeAmount = vars.flashLoanReturnAmount + actualRepayAmount - vars.repayAmount; 
                amounts = chainSwapExactOut(changeAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "7.1-no extra");

                //change the remained wBNB to USDT
                uint remain = vars.seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                vars.massProfit = amounts[amounts.length-1];
            }else{
                //case7.2 repayToken is VAI, seizedToken is wETH
                uint actualRepayAmount;
                (vars.seizedVTokenAmount, actualRepayAmount) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                vars.seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  vars.seizedVTokenAmount);

                //change part of wETH to flashLoanReturnAmount VAI
                uint changeAmount = vars.flashLoanReturnAmount + actualRepayAmount - vars.repayAmount; 
                amounts = chainSwapExactOut(changeAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "7.2-no extra");

                //change the remained wETH to USDT
                uint remain = vars.seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));

                vars.massProfit = amounts[amounts.length-1];
            }
        }else{
            revert();
        }
        repayUnderlyingToken.transfer(vars.flashLoanFrom, vars.flashLoanReturnAmount);
        emit Scenario(vars.situation, address(repayUnderlyingToken), vars.repayAmount, address(seizedUnderlyingToken), vars.flashLoanReturnAmount, vars.seizedUnderlyingAmount, vars.massProfit);
    }


    function getSeizedVToken(address _repayVToken,  address _seizedVToken, address _borrower, uint _repayAmount) internal returns (uint, uint){
        VTokenInterface seizedVToken = VTokenInterface(_seizedVToken);
        uint ok;
        uint actualRepayAmount;
        uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));

        if (isVBNB(_repayVToken)){
            IVBNB(_repayVToken).liquidateBorrow{value: _repayAmount}(_borrower, _seizedVToken); //repay BNB
        } else if (isVAI(_repayVToken)){
            (ok, actualRepayAmount) = IVAI(_repayVToken).liquidateVAI(_borrower, _repayAmount, seizedVToken);
            require(ok == 0, "liquidateBorrow error");
        }else{
            VTokenInterface repayVToken = VTokenInterface(_repayVToken);
            require(repayVToken.liquidateBorrow(_borrower, _repayAmount, seizedVToken) == 0, "liquidateBorrow error"); //repay USDT, get vUSDT
        }
        uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
        uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
        require(seizedVTokenAmount > 0, "seized VToken amount is zero");

        emit SeizedVTokenAmount(_repayAmount, actualRepayAmount, seizedVTokenAmount);
        return (seizedVTokenAmount, actualRepayAmount);
    }

    function getSeizedUnderlyingToken(address _seizedVToken, address _seizedUnderlyingToken,  uint _seizedVTokenAmount) internal returns (uint){
        uint beforeSeizedUnderlyingAmount;
        uint afterSeizedUnderlyingAmount;
        uint seizedUnderlyingAmount;

        if (isVBNB(_seizedVToken)){
            beforeSeizedUnderlyingAmount = address(this).balance;
            require(IVBNB(_seizedVToken).redeem(_seizedVTokenAmount)==0,"redeem BNB err");
            afterSeizedUnderlyingAmount = address(this).balance;
        }else{
            VTokenInterface seizedVToken = VTokenInterface(_seizedVToken);
            IERC20 seizedUnderlyingToken = IERC20(_seizedUnderlyingToken);

            beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            require(seizedVToken.redeem(_seizedVTokenAmount) == 0,"redeem error");
            afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
        }

        seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

        emit SeizedUnderlyingTokenAmount(_seizedVTokenAmount, seizedUnderlyingAmount);
        return seizedUnderlyingAmount;
    }

    function withdraw(address _token, address _to, uint _amount) onlyOwner external{
        require(_token != address(0), "token must not be zero address");
        require(_to != address(0), "to must not be zero address");
        require(_amount >0, "amount must bigger than zero");

        IERC20(_token).transfer(_to, _amount);
        emit Withdraw(_token, _to, _amount);
    }

    function withdrawETH(address payable _to, uint _amount) onlyOwner external{
        require(_to != address(0), "to must not be zero address");
        require(_amount >0, "amount must bigger than zero");

        _to.transfer(_amount);
        emit WithdrawETH(_to, _amount);
    }

    function chainSwapExactIn(uint amountIn, address[] memory path, address to) internal returns(uint[] memory amounts){
        amounts = PancakeLibrary.getAmountsOut(FACTORY, amountIn, path);
        IERC20(path[0]).transfer(PancakeLibrary.pairFor(FACTORY, path[0], path[1]), amounts[0]);
        _swap(amounts, path, to);
        return amounts;
    }


    function chainSwapExactOut(uint amountExactOut, address[] memory path, address to) internal returns(uint[] memory amounts) {
        amounts = PancakeLibrary.getAmountsIn(FACTORY, amountExactOut, path);
        IERC20(path[0]).transfer(PancakeLibrary.pairFor(FACTORY, path[0], path[1]), amounts[0]);
        _swap(amounts, path, to);
        return amounts;
    }


    // **** SWAP ****
    // requires the initial amount to have already been sent to the first pair
    function _swap(uint[] memory amounts, address[] memory path, address _to) internal {
        for (uint i; i < path.length - 1; i++) {
            (address input, address output) = (path[i], path[i + 1]);
            (address token0,) = PancakeLibrary.sortTokens(input, output);
            uint amountOut = amounts[i + 1];
            (uint amount0Out, uint amount1Out) = input == token0 ? (uint(0), amountOut) : (amountOut, uint(0));
            address to = i < path.length - 2 ? PancakeLibrary.pairFor(FACTORY, output, path[i + 2]) : _to;
            IPancakePair(PancakeLibrary.pairFor(FACTORY, input, output)).swap(
                amount0Out, amount1Out, to, new bytes(0)
            );
        }
    }


    receive() payable external{}

    function isVBNB(address _token) internal pure returns(bool){
        return (_token == vBNB);
    }

    function isStableCoin(address _token) internal pure returns (bool){
        return (_token == vBUSD || _token == vUSDT || _token == vDAI);
    }

    function isVAI(address _token) internal pure returns(bool){
        return (_token == VAI);
    }
}

