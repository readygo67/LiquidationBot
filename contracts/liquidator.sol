// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./interface.sol";
import "./PancakeLibrary.sol";


contract UniFlashSwap is IPancakeCallee {
    address private constant ComptrollerAddr = 0xfD36E2c2a6789Db23113685031d7F16329158384;
    address private constant wBNB = 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c;
    address private constant FACTORY = 0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73;
    address private constant vBNB = 0xA07c5b74C9B40447a954e1466938b865b6BBea36;
    address private constant vBUSD = 0x95c78222B3D6e262426483D42CfA53685A67Ab9D;
    address private constant vUSDT = 0xfD5840Cd36d94D7229439859C0112a4185BC0255;
    address private constant vDAI = 0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1;

    event Scenario(uint scenarioNo, address repayUnderlyingToken, uint repayAmount, address seizedUnderlyingToken, uint captureSeizedUnderlyingAmount, uint flashloanReturnAmount, uint massProfit);
    event Log(string message, uint val);
    event Log(string message, address addr);



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
    function qingsuan(uint situation,address _flashLoanFrom, address [] calldata  _path1,  address [] calldata  _path2,address [] calldata  tokens, uint _flashLoanAmount) external {
        require(situation>=1&&situation<=5,"wrong si");
        require(_flashLoanFrom != address(0), "!pair");
        //token0，token1的顺序要确定好
        address token0 = IPancakePair(_flashLoanFrom).token0();
        address token1 = IPancakePair(_flashLoanFrom).token1();
        //我们只想要一种币，看好0和1那个是我们要借的，把数设置好，另外一种币设置成0
        uint amount0Out = tokens[3] == token0 ? _flashLoanAmount : 0;
        uint amount1Out = tokens[3] == token1 ? _flashLoanAmount : 0;
        bytes memory callbackdata = abi.encode(situation,_flashLoanFrom,_path1,_path2,tokens,_flashLoanAmount);
        IPancakePair(_flashLoanFrom).swap(amount0Out, amount1Out, address(this), callbackdata);
    }

    //callback function from pair
    function pancakeCall(
        address _sender,
        uint _amount0,
        uint _amount1,
        bytes calldata _data
    ) external override {
        (uint situation,address _flashLoanFrom, address [] memory _path1,  address [] memory _path2,address [] memory tokens, uint repayAmount) = abi.decode(_data, (uint,address,address [],address [],address [],uint));

        // address token0 = IPancakePair(msg.sender).token0();
        // address token1 = IPancakePair(msg.sender).token1();

        require(msg.sender == _flashLoanFrom, "!pair");
        require(_sender == address(this), "!sender");

        uint flashLoanReturnAmount = repayAmount + ((repayAmount * 25) / 9975) + 1;

        ////path1： 卖的时候的path, seizedSymbol => repaySymbol的path
        //path2:  将seizedSymbol => USDT
        //tokens：
        // Tokens array
        // [0] - _flashLoanVToken 要去借的钱（要还给venus的）
        // [1] - _seizedVToken 可以赎回来的钱
        // [2] - _seizedTokenUnderlying 赎回来的钱的underlying
        // [3] - _flashloanTokenUnderlying 借的钱的underlying
        // [4] - target 目标账号
        VTokenInterface repayVToken = VTokenInterface(tokens[0]);
        VTokenInterface seizedVToken = VTokenInterface(tokens[1]);
        IERC20 repayUnderlyingToken = IERC20(tokens[3]);
        IERC20 seizedUnderlyingToken = IERC20(tokens[2]);
        address borrower = tokens[4];

        uint[] memory amounts;
        // uint massProfit;
        // uint beforeSeizedVTokenAmount;
        // uint afterSeizedVTokenAmount;
        // uint seizedVTokenAmount;
        // uint beforeSeizedUnderlyingAmount;
        // uint afterSeizedUnderlyingAmount;
        // uint seizedUnderlyingAmount;


        if(situation==1){
            //case1: repayToken is USDT, seizedToken is USDT
            require(_path1.length==0 && _path2.length==0,"1-patherr");
            require(isStableCoin(tokens[0]), "1-not stable coin");
            require(tokens[0] == tokens[1], "1- not same coin");

            uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"1-liquidateBorrow error "); //repay USDT, get vUSDT
            uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
            require(seizedVTokenAmount > 0,"1-seized vtoken amount is zero");

            uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            require(seizedVToken.redeem(seizedVTokenAmount) == 0,"1-redeem error");              //redeem vUSDT to USDT
            uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;
            require(seizedUnderlyingAmount > flashLoanReturnAmount, "1-no extra");

            uint massProfit = seizedUnderlyingAmount - flashLoanReturnAmount;
            emit Scenario(1, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
        }
        else if(situation==2){
            require(_path1.length==0 && _path2.length!=0,"2.1-patherr");
            if(isVBNB(tokens[0])) {
                //case2.1 repayToken is BNB, seizedToken is BNB 
                IWETH(wBNB).withdraw(repayAmount); //change the flashLoaned wBNB to BNB.

                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                IVBNB(tokens[0]).liquidateBorrow{value: repayAmount}(borrower, tokens[1]); //repay BNB，get vBNB
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"2.1-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(tokens[1]).redeem(seizedVTokenAmount)==0,"2.1-redeemerr");  //redeem vBNB to BNB
                uint afterSeizedUnderlyingAmount = address(this).balance;
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;
                require(seizedUnderlyingAmount>flashLoanReturnAmount,"2.1-no-extra");

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}(); //change BNB to wBNB

                uint remain = seizedUnderlyingAmount-flashLoanReturnAmount; //calculate how much wBNB left after return flashloan
                amounts = chainSwapExactIn(remain,_path2,address(this));  //swap the left wBNB to USDT
                uint massProfit = amounts[amounts.length-1];

                emit Scenario(2, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }
            else // not bnb
            {
                //case2.2 repayToken is wETH, seizedToken is wETH
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"2.2-liquidateBorrow error "); //repay wETH, get vETH
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"2.2-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(seizedVTokenAmount) == 0,"2.2-redeem error");          //redeem vETH to wETH
                uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;
                require(seizedUnderlyingAmount > flashLoanReturnAmount, "2.2-no extra");

                uint remain = seizedUnderlyingAmount - flashLoanReturnAmount;    //calculate how much ETH left after return flashloan
                amounts = chainSwapExactIn(remain,_path2,address(this));  //swap the left wETH to USDT
                uint massProfit = amounts[amounts.length-1];

                emit Scenario(2, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }
        }
        else if(situation==3){
            if (isVBNB(tokens[0])){
                // case3.1 seizedToken is USDT, repayToken is BNB
                IWETH(wBNB).withdraw(repayAmount); //change the flashLoaned wBNB to BNB.

                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                IVBNB(tokens[0]).liquidateBorrow{value: repayAmount}(borrower, tokens[1]); //repay BNB, get vUSDT
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"3.1-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(seizedVTokenAmount)==0,"3.1-redeemerr");  // redeem vUSDT to USDT
                uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                // change part of USDT to flashLoanReturnAmount wBNB for returning flashloan later
                amounts =  chainSwapExactOut(flashLoanReturnAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "3.1-no-extra");

                uint massProfit = seizedUnderlyingAmount - amounts[0];
                emit Scenario(3, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }else{
                // case3.2 seizedToken is USDT, repayToken is wETH
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"3.2-liquidateBorrow error "); //repay wETH, get vUSDT
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"3.2-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(seizedVTokenAmount) == 0,"3.2-redeem error");               //redeem vUSDT to USDT
                uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                // change part of USDT to flashLoanReturnAmount wETH for returning flashloan later
                amounts = chainSwapExactOut(flashLoanReturnAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "3.2-bnb-no-extra");

                uint massProfit = seizedUnderlyingAmount - amounts[0];
                emit Scenario(3, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }
        }else if(situation==4){
            if (isVBNB(tokens[1])){
                //case4.1 seizedToken is BNB, repayToken is USDT
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"4.1-liquidateBorrow error "); //repay USDT, get vBNB
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"4.1-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(tokens[1]).redeem(seizedVTokenAmount) == 0,"4.1-redeem error");    //redeem vBNB to BNB
                uint afterSeizedUnderlyingAmount = address(this).balance;
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}();  //change BNB to wBNB

                //change all wBNB to USDT
                amounts = chainSwapExactIn(seizedUnderlyingAmount, _path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > flashLoanReturnAmount, "4.1-no extra");
                uint massProfit = usdtAmount - flashLoanReturnAmount;

                emit Scenario(4, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }else{
                //case4.2 seizedToken is ETH, repayToken is USDT
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"4.2-liquidateBorrow error "); //repay USDT, get vETH
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"4.2-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(seizedVTokenAmount) == 0,"4.2-redeem error");   //redeem vETH to wETH
                uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                // change all wETH to USDT
                amounts = chainSwapExactIn(seizedUnderlyingAmount, _path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > flashLoanReturnAmount, "4.2-no extra");
                uint massProfit = usdtAmount - flashLoanReturnAmount;

                emit Scenario(4, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount,massProfit);
            }
        }else if(situation==5){
            if (isVBNB(tokens[0])){
                //case5.1 seizedToken is ETH, repayToken is BNB,
                IWETH(wBNB).withdraw(repayAmount); //change the flashLoaned wBNB to BNB.

                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                IVBNB(tokens[0]).liquidateBorrow{value: repayAmount}(borrower, tokens[1]); //repay BNB, get vETH
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"5.1-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(seizedVTokenAmount) == 0,"5.1-redeem error");   //redeem vETH to wETH
                uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                //change part of wETH to flashLoanReturnAmount wBNB
                amounts = chainSwapExactOut(flashLoanReturnAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change remain wETH to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, _path2, address(this));

                uint massProfit = amounts[amounts.length-1];
                emit Scenario(5, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }else if (isVBNB(tokens[1])){
                //case5.2 seizedToken is BNB, repayToken is ETH
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"5.2-liquidateBorrow error"); //repay ETH, get vBNB
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"5.2-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(tokens[1]).redeem(seizedVTokenAmount) == 0,"5.2-redeem error");    //redeem vBNB to BNB
                uint afterSeizedUnderlyingAmount = address(this).balance;
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}();  //change BNB to wBNB

                //change part of wBNB to flashLoanReturnAmount ETH
                amounts = chainSwapExactOut(flashLoanReturnAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change the remained wBNB to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, _path2, address(this));

                uint massProfit = amounts[amounts.length-1];
                emit Scenario(5, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }else{
                //case5.3 repayToken is wETH, seizedToken is CAKE
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                require(repayVToken.liquidateBorrow(borrower, repayAmount, seizedVToken) == 0,"5.3-liquidateBorrow error "); //repay wETH, get vCAKE
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"5.3-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                require(seizedVToken.redeem(seizedVTokenAmount) == 0,"5.3-redeem error");   //redeem vCAKE to CAKE
                uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                //change part of CAKE to flashLoanReturnAmount ETH for returning flashloan later
                amounts = chainSwapExactOut(flashLoanReturnAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "5.3-no extra");

                //change the remained CAKE to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, _path2, address(this));

                uint massProfit = amounts[amounts.length-1];
                emit Scenario(5, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }
        }else if (situation==6){
            //case6 repayToken is VAI, seizedToken is USDT
            uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            (uint ok, uint actualRepayAmount) = IVAI(tokens[0]).liquidateVAI(borrower, repayAmount, seizedVToken);
            require(ok == 0,"6-liquidateBorrow error "); //repay VAI, get vUSDT
            uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
            uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
            require(seizedVTokenAmount > 0,"5.3-seized vtoken amount is zero");

            uint beforeSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            require(seizedVToken.redeem(seizedVTokenAmount) == 0,"6-redeem error");   //redeem vUSDT to USDT
            uint afterSeizedUnderlyingAmount = seizedUnderlyingToken.balanceOf(address(this));
            uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

            // change part of USDT to flashLoanReturnAmount wVAI for returning flashloan later
            uint changeAmount = flashLoanReturnAmount + actualRepayAmount - repayAmount; 
            amounts = chainSwapExactOut(changeAmount, _path1, address(this));
            require(seizedUnderlyingAmount > amounts[0], "6-bnb-no-extra");

            uint massProfit = seizedUnderlyingAmount - amounts[0];
            emit Scenario(6, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);

        }else if (situation==7){
            //case7.1 repayToken is VAI, seizedToken is BNB
            if (isVBNB(tokens[1])){
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                (uint ok, uint actualRepayAmount) = IVAI(tokens[0]).liquidateVAI(borrower, repayAmount, seizedVToken); //repay VAI, get vBNB
                require(ok == 0,"7.1-liquidateBorrow error "); 
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"7.1-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = address(this).balance;
                require(IVBNB(tokens[1]).redeem(seizedVTokenAmount) == 0,"7.1-redeem error");    //redeem vBNB to BNB
                uint afterSeizedUnderlyingAmount = address(this).balance;
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}();  //change BNB to wBNB

                //change part of wBNB to flashLoanReturnAmount VAI
                uint changeAmount = flashLoanReturnAmount + actualRepayAmount - repayAmount; 
                amounts = chainSwapExactOut(changeAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "7.1-no extra");

                //change the remained wBNB to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, _path2, address(this));

                uint massProfit = amounts[amounts.length-1];
                emit Scenario(7, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }else{
                //case7.2 repayToken is VAI, seizedToken is wETH
                uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                (uint ok, uint actualRepayAmount) = IVAI(tokens[0]).liquidateVAI(borrower, repayAmount, seizedVToken); //repay VAI, get vETH
                require(ok == 0,"7.2-liquidateBorrow error");  
                uint afterSeizedVTokenAmount = seizedVToken.balanceOf(address(this));
                uint seizedVTokenAmount = afterSeizedVTokenAmount - beforeSeizedVTokenAmount;
                require(seizedVTokenAmount > 0,"7.2-seized vtoken amount is zero");

                uint beforeSeizedUnderlyingAmount = address(this).balance;
                require(seizedVToken.redeem(seizedVTokenAmount) == 0,"7.2-redeem error");    //redeem vETH to wETH
                uint afterSeizedUnderlyingAmount = address(this).balance;
                uint seizedUnderlyingAmount = afterSeizedUnderlyingAmount - beforeSeizedUnderlyingAmount;

                //change part of wETH to flashLoanReturnAmount VAI
                uint changeAmount = flashLoanReturnAmount + actualRepayAmount - repayAmount; 
                amounts = chainSwapExactOut(changeAmount, _path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "7.2-no extra");

                //change the remained wETH to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, _path2, address(this));

                uint massProfit = amounts[amounts.length-1];
                emit Scenario(7, address(repayUnderlyingToken), repayAmount, address(seizedUnderlyingToken), seizedUnderlyingAmount, flashLoanReturnAmount, massProfit);
            }
        }else{
            revert();
        }

        repayUnderlyingToken.transfer(_flashLoanFrom, flashLoanReturnAmount);
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
        IERC20(path[0]).transfer( PancakeLibrary.pairFor(FACTORY, path[0], path[1]), amounts[0]);
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


    function isVBNB(address _token) internal returns(bool){
        return (_token == vBNB);
    }

    function isStableCoin(address _token) internal returns (bool){
        return (_token == vBUSD || _token == vUSDT || _token == vDAI);
    }
}

