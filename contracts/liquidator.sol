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
  

    mapping(address => mapping(address => bool)) approves;

    event Scenario(uint scenarioNo, address repayUnderlyingToken, uint repayAmount, address seizedUnderlyingToken, uint flashLoanReturnAmount,uint seizedUnderlyingAmount, uint massProfit);
    event Debug1(uint, address, address[], address[], address[], uint);
    event SwapOneWBNBToUSDT(uint, uint);
    
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
        uint beforeSeizedVTokenAmount;
        uint afterSeizedVTokenAmount;
        uint seizedVTokenAmount;

        //underlyingToken
        uint beforeSeizedUnderlyingAmount;
        uint afterSeizedUnderlyingAmount;
        uint seizedUnderlyingAmount;
        
        uint massProfit;
    }

    constructor(){}

    // function swapOneWBNBToBUSD() external{
    //     uint amount = 1 ether;
    //     address[] memory path = new address[](2);
    //     path[0] = wBNB;
    //     path[1] = 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56; 
    //     chainSwapExactIn(amount, path, address(this));
    // }

      function swapOneWBNBToBUSD1() external{
        uint amount = 1 ether;
        address[] memory path = new address[](2);
        path[0] = wBNB;
        path[1] = 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56; 
        uint[] memory amounts = PancakeLibrary.getAmountsOut(FACTORY, amount, path);
        emit SwapOneWBNBToUSDT(amounts[0], amounts[1]);
       // IERC20(path[0]).transfer(PancakeLibrary.pairFor(FACTORY, path[0], path[1]), amounts[0]);
    }

    function depositOneBNB() external{
        uint amount = 1 ether;
        IWETH(wBNB).deposit{value: amount}();
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
    function qingsuan(uint _situation, address _flashLoanFrom, address [] calldata  _path1,  address [] calldata  _path2,address [] calldata  _tokens, uint _flashLoanAmount) external {
        require(_situation>=1&&_situation<=7,"wrong si");
        require(_flashLoanFrom != address(0), "!pair");

        (,,uint shortfall) = Comptroller(ComptrollerAddr).getAccountLiquidity(_tokens[4]);
        require(shortfall > 0, "shortfall must greater than zer 0");
        emit Debug1(_situation, _flashLoanFrom, _path1, _path2, _tokens, _flashLoanAmount);
        
        if (!approves[_tokens[3]][_tokens[0]]){
            IERC20(_tokens[3]).approve(_tokens[0], ~uint(0));
            approves[_tokens[3]][_tokens[0]] = true;
        }

        // uint amount = 1 ether;
        // IWETH(wBNB).deposit{value: amount}();
        // address[] memory path = new address[](2);
        // path[0] = wBNB;
        // path[1] = _tokens[3]; 
        // chainSwapExactIn(amount, path, address(this));
        


        // //token0，token1的顺序要确定好
        // address token0 = IPancakePair(_flashLoanFrom).token0();
        // address token1 = IPancakePair(_flashLoanFrom).token1();
        // //我们只想要一种币，看好0和1那个是我们要借的，把数设置好，另外一种币设置成0
        // uint amount0Out = _tokens[3] == token0 ? _flashLoanAmount : 0;
        // uint amount1Out = _tokens[3] == token1 ? _flashLoanAmount : 0;
        // bytes memory callbackdata = abi.encode(_situation,_flashLoanFrom,_path1,_path2,_tokens,_flashLoanAmount);
        // IPancakePair(_flashLoanFrom).swap(amount0Out, amount1Out, address(this), callbackdata);
    }

    //callback function from pair
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
        ////path1： 卖的时候的path, seizedSymbol => repaySymbol的path
        //path2:  将seizedSymbol => USDT
        //tokens：
        // Tokens array
        // [0] - _flashLoanVToken 要去借的钱（要还给venus的）
        // [1] - _seizedVToken 可以赎回来的钱
        // [2] - _seizedTokenUnderlying 赎回来的钱的underlying
        // [3] - _flashloanTokenUnderlying 借的钱的underlying
        // [4] - target 目标账号
        VTokenInterface repayVToken = VTokenInterface(vars.tokens[0]);
        VTokenInterface seizedVToken = VTokenInterface(vars.tokens[1]);
        IERC20 repayUnderlyingToken = IERC20(vars.tokens[3]);
        IERC20 seizedUnderlyingToken = IERC20(vars.tokens[2]);
        uint[] memory amounts;

        if(vars.situation==1){
            //case1: repayToken is USDT, seizedToken is USDT
            require(vars.path1.length==0 && vars.path2.length==0,"1-patherr");
            require(isStableCoin(vars.tokens[0]), "1-not stable coin");
            require(vars.tokens[0] == vars.tokens[1], "1- not same coin");

            vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"1-liquidateBorrow error "); //repay USDT, get vUSDT
            vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
            require(vars.seizedVTokenAmount > 0,"1-seized vtoken amount is zero");

            vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"1-redeem error");              //redeem vUSDT to USDT
            vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;
            require(vars.seizedUnderlyingAmount > vars.flashLoanReturnAmount, "1-no extra");

            vars.massProfit = vars.seizedUnderlyingAmount - vars.flashLoanReturnAmount;
        }
        else if(vars.situation==2){
            require(vars.path1.length==0 && vars.path2.length!=0,"2.1-patherr");
            if(isVBNB(vars.tokens[0])) {
                //case2.1 repayToken is BNB, seizedToken is BNB 
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.
                
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                IVBNB(vars.tokens[0]).liquidateBorrow{value: vars.repayAmount}(vars.borrower, vars.tokens[1]); //repay BNB，get vBNB
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"2.1-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(vars.tokens[1]).redeem(vars.seizedVTokenAmount)==0,"2.1-redeemerr");  //redeem vBNB to BNB
                vars.afterSeizedUnderlyingAmount = address(this).balance;
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;
                require(vars.seizedUnderlyingAmount > vars.flashLoanReturnAmount,"2.1-no-extra");

                IWETH(wBNB).deposit{value:vars.seizedUnderlyingAmount}(); //change BNB to wBNB

                uint remain = vars.seizedUnderlyingAmount-vars.flashLoanReturnAmount; //calculate how much wBNB left after return flashloan
                amounts = chainSwapExactIn(remain, vars.path2, address(this));  //swap the left wBNB to USDT
                vars.massProfit = amounts[amounts.length-1];
            }else {
                //case2.2 repayToken is wETH, seizedToken is wET
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"2.2-liquidateBorrow error "); //repay wETH, get vETH
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"2.2-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"2.2-redeem error");          //redeem vETH to wETH
                vars.afterSeizedUnderlyingAmount  = seizedUnderlyingToken.balanceOf(address(this));
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;
                require(vars.seizedUnderlyingAmount > vars.flashLoanReturnAmount, "2.2-no extra");

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

                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                IVBNB(vars.tokens[0]).liquidateBorrow{value: vars.repayAmount}(vars.borrower, vars.tokens[1]); //repay BNB, get vUSDT
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"3.1-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"3.1-redeemerr");  // redeem vUSDT to USDT
                vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

                // change part of USDT to flashLoanReturnAmount wBNB for returning flashloan later
                amounts =  chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "3.1-no-extra");

                vars.massProfit = vars.seizedUnderlyingAmount - amounts[0];
            }else{
                // case3.2 seizedToken is USDT, repayToken is wETH
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"3.2-liquidateBorrow error "); //repay wETH, get vUSDT
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"3.2-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"3.2-redeem error");               //redeem vUSDT to USDT
                vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

                // change part of USDT to flashLoanReturnAmount wETH for returning flashloan later
                amounts = chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "3.2-bnb-no-extra");

                vars.massProfit = vars.seizedUnderlyingAmount - amounts[0];
            }
        }else if(vars.situation==4){
            require(isStableCoin(vars.tokens[0]), "4-repayToken is not stable coin");
            if (isVBNB(vars.tokens[1])){
                //case4.1 seizedToken is BNB, repayToken is USDT
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"4.1-liquidateBorrow error "); //repay USDT, get vBNB
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"4.1-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(vars.tokens[1]).redeem(vars.seizedVTokenAmount) == 0,"4.1-redeem error");    //redeem vBNB to BNB
                vars.afterSeizedUnderlyingAmount = address(this).balance;
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

                IWETH(wBNB).deposit{value:vars.seizedUnderlyingAmount}();  //change BNB to wBNB

                //change all wBNB to USDT
                amounts = chainSwapExactIn(vars.seizedUnderlyingAmount, vars.path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > vars.flashLoanReturnAmount, "4.1-no extra");
                vars.massProfit = usdtAmount - vars.flashLoanReturnAmount;
            }else{
                //case4.2 seizedToken is ETH, repayToken is USDT
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"4.2-liquidateBorrow error "); //repay USDT, get vETH
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"4.2-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"4.2-redeem error");   //redeem vETH to wETH
                vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

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

                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                IVBNB(vars.tokens[0]).liquidateBorrow{value: vars.repayAmount}(vars.borrower, vars.tokens[1]); //repay BNB, get vETH
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"5.1-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"5.1-redeem error");   //redeem vETH to wETH
                vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

                //change part of wETH to flashLoanReturnAmount wBNB
                amounts = chainSwapExactOut(vars.flashLoanReturnAmount, vars.path1, address(this));
                require(vars.seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change remain wETH to USDT
                uint remain = vars.seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                vars.massProfit = amounts[amounts.length-1];

            }else if (isVBNB(vars.tokens[1])){
                //case5.2 seizedToken is BNB, repayToken is ETH
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"5.2-liquidateBorrow error"); //repay ETH, get vBNB
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"5.2-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(vars.tokens[1]).redeem(vars.seizedVTokenAmount) == 0,"5.2-redeem error");    //redeem vBNB to BNB
                vars.afterSeizedUnderlyingAmount = address(this).balance;
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

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
                vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(vars.borrower, vars.repayAmount, seizedVToken) == 0,"5.3-liquidateBorrow error "); //repay wETH, get vCAKE
                vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"5.3-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"5.3-redeem error");   //redeem vCAKE to CAKE
                vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

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
         
            vars.beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            (uint ok, uint actualRepayAmount) = IVAI(vars.tokens[0]).liquidateVAI(vars.borrower, vars.repayAmount, seizedVToken);
            require(ok == 0,"6-liquidateBorrow error "); //repay VAI, get vUSDT
            vars.afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
            require(vars.seizedVTokenAmount > 0,"6-seized vtoken amount is zero");
            
            vars.beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"6-redeem error");   //redeem vUSDT to USDT
            vars.afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;
            
            // change part of USDT to flashLoanReturnAmount VAI for returning flashloan later
            uint changeAmount = vars.flashLoanReturnAmount + actualRepayAmount - vars.repayAmount;
            chainSwapExactOut(changeAmount, vars.path1, address(this));
            require(vars.seizedUnderlyingAmount > amounts[0], "6-noextra");

            vars.massProfit = vars.seizedUnderlyingAmount - amounts[0];
        }else if (vars.situation==7){
            //case7.1 repayToken is VAI, seizedToken is BNB
            if (isVBNB(vars.tokens[1])){

                vars.beforeSeizedUnderlyingAmount = seizedVToken.balanceOf(address(this));
                (uint ok, uint actualRepayAmount) = IVAI(vars.tokens[0]).liquidateVAI(vars.borrower, vars.repayAmount, seizedVToken); //repay VAI, get vBNB
                require(ok == 0,"7.1-liquidateBorrow error "); 
                vars.afterSeizedUnderlyingAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"7.1-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(vars.tokens[1]).redeem(vars.seizedVTokenAmount) == 0,"7.1-redeem error");    //redeem vBNB to BNB
                vars.afterSeizedUnderlyingAmount = address(this).balance;
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

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
                vars.beforeSeizedUnderlyingAmount = seizedVToken.balanceOf(address(this));
                (uint ok, uint actualRepayAmount) = IVAI(vars.tokens[0]).liquidateVAI(vars.borrower, vars.repayAmount, seizedVToken); //repay VAI, get vETH
                require(ok == 0,"7.2-liquidateBorrow error");  
                vars.afterSeizedUnderlyingAmount = seizedVToken.balanceOf(address(this));
                vars.seizedVTokenAmount = vars.afterSeizedVTokenAmount - vars.beforeSeizedVTokenAmount;
                require(vars.seizedVTokenAmount > 0,"7.2-seized vtoken amount is zero");

                vars.beforeSeizedUnderlyingAmount = address(this).balance;
                require(seizedVToken.redeem(vars.seizedVTokenAmount) == 0,"7.2-redeem error");    //redeem vETH to wETH
                vars.afterSeizedUnderlyingAmount = address(this).balance;
                vars.seizedUnderlyingAmount = vars.afterSeizedUnderlyingAmount - vars.beforeSeizedUnderlyingAmount;

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


    function chainSwapExactIn(uint amountIn, address[] memory path, address to) internal returns(uint[] memory amounts){
        amounts = PancakeLibrary.getAmountsOut(FACTORY, amountIn, path);
        //把path0的钱撞到pair里
        // TransferHelper.safeTransferFrom(
        //     path[0], msg.sender, PancakeLibrary.pairFor(factory, path[0], path[1]), amounts[0]
        // );
        IERC20(path[0]).transfer(PancakeLibrary.pairFor(FACTORY, path[0], path[1]), amounts[0]);
        _swap(amounts, path, to);
        return amounts;
    }


    function chainSwapExactOut(uint amountExactOut, address[] memory path, address to) internal returns(uint[] memory amounts) {
        amounts = PancakeLibrary.getAmountsIn(FACTORY, amountExactOut, path);
        //把path0的钱撞到pair里
        // TransferHelper.safeTransferFrom(
        //     path[0], msg.sender, PancakeLibrary.pairFor(factory, path[0], path[1]), amounts[0]
        // );
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

    function withdraw(address _token, uint _amount) onlyOwner external{
        require(_token != address(0), "address must not be zero");
        require(_amount >0, "amount must bigger than zero");
        IERC20(_token).transfer(msg.sender, _amount);
    }

    receive() payable external{}

    function isVBNB(address _token) internal pure returns(bool){
        return (_token == vBNB);
    }

    function isStableCoin(address _token) internal pure returns (bool){
        return (_token == vBUSD || _token == vUSDT || _token == vDAI);
    }

}

