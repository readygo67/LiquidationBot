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
    address private constant VAIController = 0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE;
    address private constant USDT = 0x55d398326f99059fF775485246999027B3197955;
    uint private constant MAXUINT32 = ~uint(0);

    mapping(address => mapping(address => bool)) approves;

    event Scenario(uint scenarioNo, address repayUnderlyingToken, uint repayAmount, address seizedUnderlyingToken, uint flashLoanReturnAmount,uint seizedUnderlyingAmount, uint massProfit);
    event SeizedVTokenAmount(uint, uint, uint);
    event SeizedUnderlyingTokenAmount(uint, uint);
    event Withdraw(address indexed, address indexed, uint);
    event WithdrawETH(address indexed, uint);

    struct LocalVars {
        uint situation;
        address flashLoanFrom;
        address[] path1;
        address[] path2;
        address[] tokens;
        uint repayAmount;
    }

    function swapOneBNBToFlashLoandUnderlyingToken(address _flashLoanUnderlyingToken) onlyOwner public{
        uint amount = 1 ether;
         IWETH(wBNB).deposit{value: amount}();

        address[] memory path = new address[](2);
        path[0] = wBNB;
        path[1] = _flashLoanUnderlyingToken; 
        chainSwapExactIn(amount, path, address(this));
    }

    function mintedVAIS(address _account) public view returns(uint) {
        return Comptroller(ComptrollerAddr).mintedVAIs(_account);
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
    case1 gas: 752659
       0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE
       ["0xfD5840Cd36d94D7229439859C0112a4185BC0255","0xfD5840Cd36d94D7229439859C0112a4185BC0255","0x55d398326f99059fF775485246999027B3197955","0x55d398326f99059fF775485246999027B3197955","0x564EE8bF0bA977A1ccc92fe3D683AbF4569c9f5E"]
       21034205123917372652

    case2.1(TODO)




    case2.2 gas:1382255
        height15379782, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, repaySmbol:vETH, flashLoanFrom:0x74E4716E431f45807DCF19f284c7aA99F18a4fbc, repayAddress:0xf508fCD89b8bd15579dc79A6827cB4686A3592c8, repayValue:27760775212595059202.7, repayAmount:9834830202498401 seizedSymbol:vETH, seizedAddress:0xf508fCD89b8bd15579dc79A6827cB4686A3592c8, seizedCTokenAmount:53542165, seizedUnderlyingTokenAmount:10818313118937201.583675475128906, seizedUnderlyingTokenValue:30536852440824038910.2407636463629662
        calculateSeizedTokenAmount case2: seizedSymbol == repaySymbol and symbol is not stable coin, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, symbol:vETH, seizedAmount:10818313118937201.583675475128906, returnAmout:9859478899747771, usdtAmount:2702645379426654958, gasFee:2425666800000000000, profit:0.2786001639516656
        case2, profitable liquidation catched:&{0xFAbE4C180b6eDad32eA0Cf56587c54417189e422 0.9512268623785098 15379543 0001-01-01 00:00:00 +0000 UTC}, profit:0.2786001639516656

        0x74E4716E431f45807DCF19f284c7aA99F18a4fbc
        path1: ["0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x55d398326f99059fF775485246999027B3197955"] 
        path2: ["0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x55d398326f99059fF775485246999027B3197955"]
        tokens:["0xf508fCD89b8bd15579dc79A6827cB4686A3592c8","0xf508fCD89b8bd15579dc79A6827cB4686A3592c8","0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x2170ed0880ac9a755fd29b2688956bd959f933f8","0xFAbE4C180b6eDad32eA0Cf56587c54417189e422"]
        9834830202498401

    case3.1(TODO)




    case3.2 gas: 1069882
        calculateSeizedTokenAmount case3: seizedSymbol != repaySymbol and seizedSymbol stable coin
        height:15391594, account:0xF2455A4c6fcC6F41f59222F4244AFdDC85ff1Ed7, repaySymbol:vUSDC, repayUnderlyingAmount:27875946916574608303, seizedSymbol:vBUSD, seizedVTokenAmount:143329844655, seizedUnderlyingAmount:30650960029190065251.1254477794787309, seizedValue:30645088960400314010.8868557044583629, flashLoanReturnAmout:27945811445187577219, remain:2606405038873928056.1254477794787309, gasFee:2971205250000000000, profit:-0.3652994575856481
        flashLoanFrom:0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b, path1:[0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d], path2:<nil>, addresses:[0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8 0x95c78222B3D6e262426483D42CfA53685A67Ab9D 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0xF2455A4c6fcC6F41f59222F4244AFdDC85ff1Ed7]

        flashLoanFrom:0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b
        path1:["0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56","0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"]
        path2:["0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56","0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"]
        tokens:["0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8","0x95c78222B3D6e262426483D42CfA53685A67Ab9D","0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56","0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d","0xF2455A4c6fcC6F41f59222F4244AFdDC85ff1Ed7"]

    case4.1, gas 985517 => 807554
        calculateSeizedTokenAmount case4: seizedSymbol is not stable coin, repaySymbol is stable coin
        height:15416065, account:0x408C15Dd98A3F4Bb416Fd9E286cAc9a511894Bd3, repaySymbol:vBUSD, repayUnderlyingAmount:10060116064070759125, seizedSymbol:vBNB, seizedVTokenAmount:134837315, seizedUnderlyingAmount:29056432153150154.412562002028852, seizedValue:11079194334850431357.3863613239996445, flashLoanReturnAmout:10085329387539608136, remain:938946622105055189, gasFee:3812992000000000000, profit:-2.8740453778949448
        flashLoanFrom:0x7EFaEf62fDdCCa950418312c6C91Aef321375A00, path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56], path2:<nil>, addresses:[0x95c78222B3D6e262426483D42CfA53685A67Ab9D 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 0x408C15Dd98A3F4Bb416Fd9E286cAc9a511894Bd3]

        flashLoanFrom:0x7EFaEf62fDdCCa950418312c6C91Aef321375A00 
        path1:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"]
        path2:<nil>
        addresses:["0x95c78222B3D6e262426483D42CfA53685A67Ab9D", "0xA07c5b74C9B40447a954e1466938b865b6BBea36", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56", "0x408C15Dd98A3F4Bb416Fd9E286cAc9a511894Bd3"]


    case4.2 gas:1666282 => 1181926 => 1370176
        calculateSeizedTokenAmount case4: seizedSymbol is not stable coin, repaySymbol is stable coin
        height:15415790, account:0x2DdD4C35C2B51b75a2B259C6b102c82EBc67881B, repaySymbol:vDAI, repayUnderlyingAmount:150399177484357810, seizedSymbol:vETH, seizedVTokenAmount:309232, seizedUnderlyingAmount:62481733237641.9379836833174916, seizedValue:165940861584526588.155205665093716, flashLoanReturnAmout:150776117778804822, remain:13868787239816813, gasFee:3817500000000000000, profit:-3.8036303667641616
        flashLoanFrom:0xc7c3cCCE4FA25700fD5574DA7E200ae28BBd36A3, path1:[0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3], path2:<nil>, addresses:[0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1 0xf508fCD89b8bd15579dc79A6827cB4686A3592c8 0x2170Ed0880ac9A755fd29B2688956BD959F933F8 0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3 0x2DdD4C35C2B51b75a2B259C6b102c82EBc67881B]

        flashLoanFrom:0xc7c3cCCE4FA25700fD5574DA7E200ae28BBd36A3
        path1:["0x2170Ed0880ac9A755fd29B2688956BD959F933F8", "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3"] 
        path2:<nil>, 
        addresses:["0x334b3eCB4DCa3593BCCC3c7EBD1A1C1d1780FBF1", "0xf508fCD89b8bd15579dc79A6827cB4686A3592c8", "0x2170Ed0880ac9A755fd29B2688956BD959F933F8", "0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3", "0x2DdD4C35C2B51b75a2B259C6b102c82EBc67881B"]
        150399177484357810
    
    
    case5.1 gas:1208500 => 974520
        calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
        height:15415777, account:0x1611561c252dC43828236697ab90F278C5e674eC, repaySymbol:vBNB, repayUnderlyingAmount:50087024857209646, seizedSymbol:vBTC, seizedVTokenAmount:2680833, seizedUnderlyingAmount:541650827784433.5304300612567128, seizedValue:21024545203015680898.4435201209618967, flashLoanReturnAmout:50212556247829219, remain:47721527636533, gasFee:2863125000000000000, profit:-1.0170406768993241
        flashLoanFrom:0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16, path1:[0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c], path2:[0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c 0x55d398326f99059fF775485246999027B3197955], addresses:[0xA07c5b74C9B40447a954e1466938b865b6BBea36 0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B 0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x1611561c252dC43828236697ab90F278C5e674eC]

        flashLoanFrom:0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16
        path1:["0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"]
        path2:["0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c", "0x55d398326f99059fF775485246999027B3197955"]
        addresses:["0xA07c5b74C9B40447a954e1466938b865b6BBea36", "0x882C173bC7Ff3b7786CA16dfeD3DFFfb9Ee7847B", "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x1611561c252dC43828236697ab90F278C5e674eC"]
        50087024857209646

    case5.2 gas:927990 => 818503 => 956613
        calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
        height:15415559, account:0x5CA18142b3Aa0E4e580ec939fe3761af37395262, repaySymbol:vTUSD, repayUnderlyingAmount:5012979382087345840, seizedSymbol:vBNB, seizedVTokenAmount:67209325, seizedUnderlyingAmount:14483051888262770.9963578757692161, seizedValue:5504860657675716182.0207406384930878, flashLoanReturnAmout:5025543240187815374, remain:1606269562758565, gasFee:2850673687500000000, profit:-2.2406430009646194
        flashLoanFrom:0x2E28b9B74D6d99D4697e913b82B41ef1CAC51c6C, path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x14016E85a25aeb13065688cAFB43044C2ef86784], path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955], addresses:[0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x14016E85a25aeb13065688cAFB43044C2ef86784 0x5CA18142b3Aa0E4e580ec939fe3761af37395262]

        flashLoanFrom:0x2E28b9B74D6d99D4697e913b82B41ef1CAC51c6C
        path1:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x14016E85a25aeb13065688cAFB43044C2ef86784"]
        path2:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x55d398326f99059fF775485246999027B3197955"]
        addresses:["0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3", "0xA07c5b74C9B40447a954e1466938b865b6BBea36", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x14016E85a25aeb13065688cAFB43044C2ef86784", "0x5CA18142b3Aa0E4e580ec939fe3761af37395262"]

    case5.3 gas:1189122 => 998937 => 1212872
        calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
        height:15415786, account:0x2eB71e5335d5328e76fa0755Db27E184Be834D31, repaySymbol:vUSDC, repayUnderlyingAmount:1257422920232379755, seizedSymbol:vCAKE, seizedVTokenAmount:797776351, seizedUnderlyingAmount:186477629725528706.5014612928402019, seizedValue:1383477534933697473.5343413315814579, flashLoanReturnAmout:1260574356122686470, remain:15530062403094129, gasFee:2863125000000000000, profit:-2.7482294784463128
        flashLoanFrom:0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b, path1:[0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d], path2:[0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x55d398326f99059fF775485246999027B3197955], addresses:[0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8 0x86aC3974e2BD0d60825230fa6F355fF11409df5c 0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d 0x2eB71e5335d5328e76fa0755Db27E184Be834D31]

        flashLoanFrom:0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b
        path1:["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82", "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"]
        path2:["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82", "0x55d398326f99059fF775485246999027B3197955"]
        addresses:["0xecA88125a5ADbe82614ffC12D0DB554E2e2867C8", "0x86aC3974e2BD0d60825230fa6F355fF11409df5c", "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82", "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d", "0x2eB71e5335d5328e76fa0755Db27E184Be834D31"]

    case6(TODO) 



    case7.1 gas:1516478 =>1036778
        calculateSeizedTokenAmount case7: repaySymbol is VAI and seizedSymbol is not stable coin
        height:15413782, account:0x614146018042D47Dcde01A9400A8d14343047b67, repaySymbol:VAI, repayUnderlyingAmount:16699314059311285959, seizedSymbol:vBNB, seizedVTokenAmount:223898023, seizedUnderlyingAmount:48247431948231311.458105709752904, seizedValue:18322461643261504720.2892346727612366, flashLoanReturnAmout:16741166976753168865, remain:8540375333930007, gasFee:2848202625000000000, profit:0.3943119508401805
        flashLoanFrom:0x133ee93FE93320e1182923E1a640912eDE17C90C, path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7], path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955], addresses:[0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7 0x614146018042D47Dcde01A9400A8d14343047b67]
        
        flashLoanFrom: 0x133ee93FE93320e1182923E1a640912eDE17C90C
        path1:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7"]
        path2:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x55d398326f99059fF775485246999027B3197955"]
        addresses:["0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE", "0xA07c5b74C9B40447a954e1466938b865b6BBea36", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7", "0x614146018042D47Dcde01A9400A8d14343047b67"]
        16699314059311285959

    case7.2 gas:2112126 ==> 1371980 => 2090456
        calculateSeizedTokenAmount case7: repaySymbol is VAI and seizedSymbol is not stable coin
        height:15413870, account:0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7, repaySymbol:VAI, repayUnderlyingAmount:1539941852601276, seizedSymbol:vXVS, seizedVTokenAmount:904469, seizedUnderlyingAmount:181895892146378.5382049533939062, seizedValue:1694360235343516.0833791408642363, flashLoanReturnAmout:1543801355991255, remain:122091524525352, gasFee:2858175000000000000, profit:-2.8571194908603029
        flashLoanFrom:0x133ee93FE93320e1182923E1a640912eDE17C90C, path1:[0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7], path2:[0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x55d398326f99059fF775485246999027B3197955], addresses:[0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE 0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D 0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63 0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7 0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7]

        flashLoanFrom:0x133ee93FE93320e1182923E1a640912eDE17C90C, 
        path1:["0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63","0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7"] 
        path2:["0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63", "0x55d398326f99059fF775485246999027B3197955"]
        addresses:["0x004065D34C6b18cE4370ced1CeBDE94865DbFAFE", "0x151B1e2635A717bcDc836ECd6FbB62B674FE3E1D", "0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63", "0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7", "0x9F7A5885051fB71c4D2a7aB4203446FaCdF65BF7"]
        1539941852601276

    */
    function qingsuan(uint _situation, address _flashLoanFrom, address [] calldata  _path1,  address [] calldata  _path2,address [] calldata  _tokens, uint _flashLoanAmount) external {
       {
           uint allowed = Comptroller(ComptrollerAddr).liquidateBorrowAllowed(_tokens[0], _tokens[1], address(this), _tokens[4], _flashLoanAmount);
           require(allowed == 0,"sanity check fail");
        //   (,,uint shortfall) = Comptroller(ComptrollerAddr).getAccountLiquidity(_tokens[4]);
        //   require(shortfall > 0, "shortfall must greater than zer 0");

        //    uint closeFactor = Comptroller(ComptrollerAddr).closeFactorMantissa();
        //    uint borrowAmount;

        //    if (_situation < 6){
        //        borrowAmount = VTokenInterface(_tokens[0]).borrowBalanceStored(_tokens[4]);
        //    }else{
        //        borrowAmount = Comptroller(ComptrollerAddr).mintedVAIs(_tokens[4]);
        //    }
        //    uint maxAllowedRepayAmount = (borrowAmount * closeFactor)/ 1e18;
        //    require(_flashLoanAmount <= maxAllowedRepayAmount, "flashloanAmount must be less than maxAllowedRepayAmount");
       }

        if (!approves[_tokens[3]][_tokens[0]]){
            IERC20(_tokens[3]).approve(_tokens[0], MAXUINT32);
            approves[_tokens[3]][_tokens[0]] = true;
        }

        //token0，token1的顺序要确定好
        address token0 = IPancakePair(_flashLoanFrom).token0();
        address token1 = IPancakePair(_flashLoanFrom).token1();
        //我们只想要一种币，看好0和1那个是我们要借的，把数设置好，另外一种币设置成0
        uint amount0Out = _tokens[3] == token0 ? _flashLoanAmount : 0;
        uint amount1Out = _tokens[3] == token1 ? _flashLoanAmount : 0;
        bytes memory callbackdata = abi.encode(_situation,_flashLoanFrom,_path1,_path2,_tokens,_flashLoanAmount);
        IPancakePair(_flashLoanFrom).swap(amount0Out, amount1Out, address(this), callbackdata);
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

        IERC20 repayUnderlyingToken = IERC20(vars.tokens[3]);
        IERC20 seizedUnderlyingToken = IERC20(vars.tokens[2]);
        uint flashLoanReturnAmount = vars.repayAmount + ((vars.repayAmount * 25) / 9975) + 1;
        uint seizedVTokenAmount;
        uint seizedUnderlyingAmount;
        uint massProfit;

        uint[] memory amounts;

        if(vars.situation==1){
            //case1: repayToken is USDT, seizedToken is USDT
            // require(vars.path1.length==0 && vars.path2.length==0,"1-patherr");
            require(isStableCoin(vars.tokens[0]), "1-not stable coin");
            require(vars.tokens[0] == vars.tokens[1], "1- not same coin");

            (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
            seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);
            require(seizedUnderlyingAmount > flashLoanReturnAmount, "no extra");
            massProfit = seizedUnderlyingAmount - flashLoanReturnAmount;
        }
        else if(vars.situation==2){
            if(isVBNB(vars.tokens[0])) {
                //case2.1 repayToken is BNB, seizedToken is BNB
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.

                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);
                require(seizedUnderlyingAmount > flashLoanReturnAmount,"2.1-no-extra");

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}(); //change BNB to wBNB

                uint remain = seizedUnderlyingAmount-flashLoanReturnAmount; //calculate how much wBNB left after return flashloan
                amounts = chainSwapExactIn(remain, vars.path2, address(this));  //swap the left wBNB to USDT
                massProfit = amounts[amounts.length-1];
            }else {
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);
                require(seizedUnderlyingAmount > flashLoanReturnAmount,"2.2-no-extra");

                uint remain = seizedUnderlyingAmount - flashLoanReturnAmount;    //calculate how much ETH left after return flashloan
                amounts = chainSwapExactIn(remain,vars.path2,address(this));  //swap the left wETH to USDT
                massProfit = amounts[amounts.length-1];
            }
        }
        else if(vars.situation==3){
            require(isStableCoin(vars.tokens[1]), "3-seized token is not stable coin");
            if (isVBNB(vars.tokens[0])){
                // case3.1 seizedToken is USDT, repayToken is BNB
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.

                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                // change part of USDT to flashLoanReturnAmount wBNB for returning flashloan later
                amounts =  chainSwapExactOut(flashLoanReturnAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "3.1-no-extra");
                massProfit = seizedUnderlyingAmount - amounts[0];
            }else{
                // case3.2 seizedToken is USDT, repayToken is wETH
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                // change part of USDT to flashLoanReturnAmount wETH for returning flashloan later
                amounts = chainSwapExactOut(flashLoanReturnAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "3.2-bnb-no-extra");
                massProfit = seizedUnderlyingAmount - amounts[0];
            }
        }else if(vars.situation==4){
            require(isStableCoin(vars.tokens[0]), "4-repayToken is not stable coin");
            if (isVBNB(vars.tokens[1])){
                //case4.1 seizedToken is BNB, repayToken is USDT
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}();  //change BNB to wBNB

                //change all wBNB to USDT
                amounts = chainSwapExactIn(seizedUnderlyingAmount, vars.path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > flashLoanReturnAmount, "4.1-no extra");
                massProfit = usdtAmount - flashLoanReturnAmount;
            }else{
                //case4.2 seizedToken is ETH, repayToken is USDT
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                // change all wETH to USDT
                amounts = chainSwapExactIn(seizedUnderlyingAmount, vars.path1, address(this));
                uint usdtAmount = amounts[amounts.length-1];
                require(usdtAmount > flashLoanReturnAmount, "4.2-no extra");
                massProfit = usdtAmount - flashLoanReturnAmount;
            }
        }else if(vars.situation==5){
            if (isVBNB(vars.tokens[0])){
                //case5.1 seizedToken is ETH, repayToken is BNB,
                IWETH(wBNB).withdraw(vars.repayAmount); //change the flashLoaned wBNB to BNB.
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                //change part of wETH to flashLoanReturnAmount wBNB
                amounts = chainSwapExactOut(flashLoanReturnAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change remain wETH to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                massProfit = amounts[amounts.length-1];

            }else if (isVBNB(vars.tokens[1])){
                //case5.2 seizedToken is BNB, repayToken is ETH
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}();  //change BNB to wBNB

                //change part of wBNB to flashLoanReturnAmount ETH
                amounts = chainSwapExactOut(flashLoanReturnAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "5.1-no extra");

                //change the remained wBNB to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                massProfit = amounts[amounts.length-1];
            }else{
                //case5.3 repayToken is wETH, seizedToken is CAKE
                (seizedVTokenAmount, ) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                //change part of CAKE to flashLoanReturnAmount ETH for returning flashloan later
                amounts = chainSwapExactOut(flashLoanReturnAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "5.3-no extra");

                //change the remained CAKE to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                massProfit = amounts[amounts.length-1];
            }
        }else if (vars.situation==6){
            require(isStableCoin(vars.tokens[1]), "6-seizedToken is not stable coin");
            //case6 repayToken is VAI, seizedToken is USDT
            uint actualRepayAmount;
            (seizedVTokenAmount, actualRepayAmount) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
            seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

            // change part of USDT to flashLoanReturnAmount VAI for returning flashloan later
            uint changeAmount = flashLoanReturnAmount + actualRepayAmount - vars.repayAmount;
            chainSwapExactOut(changeAmount, vars.path1, address(this));
            require(seizedUnderlyingAmount > amounts[0], "6-noextra");
            massProfit = seizedUnderlyingAmount - amounts[0];
        }else if (vars.situation==7){
            //case7.1 repayToken is VAI, seizedToken is BNB
            if (isVBNB(vars.tokens[1])){
                uint actualRepayAmount;
                (seizedVTokenAmount, actualRepayAmount) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);
                IWETH(wBNB).deposit{value:seizedUnderlyingAmount}();  //change BNB to wBNB

                //change part of wBNB to flashLoanReturnAmount VAI
                uint changeAmount = flashLoanReturnAmount + actualRepayAmount - vars.repayAmount;
                amounts = chainSwapExactOut(changeAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "7.1-no extra");

                //change the remained wBNB to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                massProfit = amounts[amounts.length-1];
            }else{
                //case7.2 repayToken is VAI, seizedToken is wETH
                uint actualRepayAmount;
                (seizedVTokenAmount, actualRepayAmount) = getSeizedVToken(vars.tokens[0], vars.tokens[1], vars.tokens[4], vars.repayAmount);
                seizedUnderlyingAmount = getSeizedUnderlyingToken(vars.tokens[1], vars.tokens[2],  seizedVTokenAmount);

                //change part of wETH to flashLoanReturnAmount VAI
                uint changeAmount = flashLoanReturnAmount + actualRepayAmount - vars.repayAmount;
                amounts = chainSwapExactOut(changeAmount, vars.path1, address(this));
                require(seizedUnderlyingAmount > amounts[0], "7.2-no extra");

                //change the remained wETH to USDT
                uint remain = seizedUnderlyingAmount - amounts[0];
                amounts = chainSwapExactIn(remain, vars.path2, address(this));
                massProfit = amounts[amounts.length-1];
            }
        }else{
            revert();
        }

        repayUnderlyingToken.transfer(vars.flashLoanFrom, flashLoanReturnAmount);
        emit Scenario(vars.situation, address(repayUnderlyingToken), vars.repayAmount, address(seizedUnderlyingToken), flashLoanReturnAmount, seizedUnderlyingAmount, massProfit);
    }


    function getSeizedVToken(address _repayVToken,  address _seizedVToken, address _borrower, uint _repayAmount) internal returns (uint, uint){
        VTokenInterface seizedVToken = VTokenInterface(_seizedVToken);
        uint ok;
        uint actualRepayAmount;
        uint beforeSeizedVTokenAmount = seizedVToken.balanceOf(address(this));

        if (isVBNB(_repayVToken)){
            IVBNB(_repayVToken).liquidateBorrow{value: _repayAmount}(_borrower, _seizedVToken); //repay BNB
        } else if (isVAIController(_repayVToken)){
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

    function isVAIController(address _addr) internal pure returns(bool){
        return (_addr == VAIController);
    }
}

