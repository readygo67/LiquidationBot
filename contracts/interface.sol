// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

interface IERC20 {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);

    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
    function totalSupply() external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint value) external returns (bool);
    function transfer(address to, uint value) external returns (bool);
    function transferFrom(address from, address to, uint value) external returns (bool);
}

interface IWETH {
    function deposit() external payable;
    function transfer(address to, uint value) external returns (bool);
    function withdraw(uint) external;
}


interface VTokenInterface {


    /*** Market Events ***/

    /**
     * @notice Event emitted when interest is accrued
     */
    event AccrueInterest(uint cashPrior, uint interestAccumulated, uint borrowIndex, uint totalBorrows);

    /**
     * @notice Event emitted when tokens are minted
     */
    event Mint(address minter, uint mintAmount, uint mintTokens);

    /**
     * @notice Event emitted when tokens are minted behalf by payer to receiver
     */
    event MintBehalf(address payer, address receiver, uint mintAmount, uint mintTokens);

    /**
     * @notice Event emitted when tokens are redeemed
     */
    event Redeem(address redeemer, uint redeemAmount, uint redeemTokens);

    /**
     * @notice Event emitted when tokens are redeemed and fee are transferred
     */
    event RedeemFee(address redeemer, uint feeAmount, uint redeemTokens);

    /**
     * @notice Event emitted when underlying is borrowed
     */
    event Borrow(address borrower, uint borrowAmount, uint accountBorrows, uint totalBorrows);

    /**
     * @notice Event emitted when a borrow is repaid
     */
    event RepayBorrow(address payer, address borrower, uint repayAmount, uint accountBorrows, uint totalBorrows);

    /**
     * @notice Event emitted when a borrow is liquidated
     */
    event LiquidateBorrow(address liquidator, address borrower, uint repayAmount, address vTokenCollateral, uint seizeTokens);


    /*** Admin Events ***/

    /**
     * @notice Event emitted when pendingAdmin is changed
     */
    event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin);

    /**
     * @notice Event emitted when pendingAdmin is accepted, which means admin is updated
     */
    event NewAdmin(address oldAdmin, address newAdmin);



    /**
     * @notice Event emitted when the reserve factor is changed
     */
    event NewReserveFactor(uint oldReserveFactorMantissa, uint newReserveFactorMantissa);

    /**
     * @notice Event emitted when the reserves are added
     */
    event ReservesAdded(address benefactor, uint addAmount, uint newTotalReserves);

    /**
     * @notice Event emitted when the reserves are reduced
     */
    event ReservesReduced(address admin, uint reduceAmount, uint newTotalReserves);

    /**
     * @notice EIP20 Transfer event
     */
    event Transfer(address indexed from, address indexed to, uint amount);

    /**
     * @notice EIP20 Approval event
     */
    event Approval(address indexed owner, address indexed spender, uint amount);

    /**
     * @notice Failure event
     */
    event Failure(uint error, uint info, uint detail);

    function underlying() external view returns (address);

    /*** User Interface ***/

    function transfer(address dst, uint amount) external  returns (bool);
    function transferFrom(address src, address dst, uint amount) external  returns (bool);
    function approve(address spender, uint amount) external  returns (bool);
    function allowance(address owner, address spender) external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function balanceOfUnderlying(address owner) external returns (uint);
    function getAccountSnapshot(address account) external view returns (uint, uint, uint, uint);
    function borrowRatePerBlock() external view returns (uint);

    function supplyRatePerBlock() external view returns (uint);
    function totalBorrowsCurrent() external returns (uint);
    function borrowBalanceCurrent(address account) external returns (uint);
    function getCash() external view returns (uint);
    function seize(address liquidator, address borrower, uint seizeTokens) external returns (uint);

    /*** User Interface ***/

    function mint(uint mintAmount) external returns (uint);
    function redeem(uint redeemTokens) external returns (uint);
    function redeemUnderlying(uint redeemAmount) external returns (uint);
    function borrow(uint borrowAmount) external returns (uint);
    function repayBorrow(uint repayAmount) external returns (uint);
    function repayBorrowBehalf(address borrower, uint repayAmount) external returns (uint);
    function liquidateBorrow(address borrower, uint repayAmount, VTokenInterface vTokenCollateral) external returns (uint);


    /*** Admin Functions ***/

    function _addReserves(uint addAmount) external returns (uint);


}






interface IPancakeCallee {
    function pancakeCall(address sender, uint amount0, uint amount1, bytes calldata data) external;
}


interface IPancakePair {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);

    function name() external pure returns (string memory);
    function symbol() external pure returns (string memory);
    function decimals() external pure returns (uint8);
    function totalSupply() external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint value) external returns (bool);
    function transfer(address to, uint value) external returns (bool);
    function transferFrom(address from, address to, uint value) external returns (bool);

    function DOMAIN_SEPARATOR() external view returns (bytes32);
    function PERMIT_TYPEHASH() external pure returns (bytes32);
    function nonces(address owner) external view returns (uint);

    function permit(address owner, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s) external;

    event Mint(address indexed sender, uint amount0, uint amount1);
    event Burn(address indexed sender, uint amount0, uint amount1, address indexed to);
    event Swap(
        address indexed sender,
        uint amount0In,
        uint amount1In,
        uint amount0Out,
        uint amount1Out,
        address indexed to
    );
    
    event Sync(uint112 reserve0, uint112 reserve1);

    function MINIMUM_LIQUIDITY() external pure returns (uint);
    function factory() external view returns (address);
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    function price0CumulativeLast() external view returns (uint);
    function price1CumulativeLast() external view returns (uint);
    function kLast() external view returns (uint);

    function mint(address to) external returns (uint liquidity);
    function burn(address to) external returns (uint amount0, uint amount1);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function skim(address to) external;
    function sync() external;

    function initialize(address, address) external;
}



interface IPancakeFactory {
    event PairCreated(address indexed token0, address indexed token1, address pair, uint);

    function feeTo() external view returns (address);
    function feeToSetter() external view returns (address);

    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);

    function createPair(address tokenA, address tokenB) external returns (address pair);

    function setFeeTo(address) external;
    function setFeeToSetter(address) external;
}

interface Comptroller {
  function accountAssets ( address, uint256 ) external view returns ( address );
  function admin (  ) external view returns ( address );
  function allMarkets ( uint256 ) external view returns ( address );
  function borrowAllowed ( address vToken, address borrower, uint256 borrowAmount ) external returns ( uint256 );
  function borrowCapGuardian (  ) external view returns ( address );
  function borrowCaps ( address ) external view returns ( uint256 );
  function borrowGuardianPaused ( address ) external view returns ( bool );
  function borrowVerify ( address vToken, address borrower, uint256 borrowAmount ) external;
  function checkMembership ( address account, address vToken ) external view returns ( bool );
  function claimVenus ( address holder, address[] memory vTokens ) external;
  function claimVenus ( address holder ) external;
  function claimVenus ( address[] memory holders, address[] memory vTokens, bool borrowers, bool suppliers ) external;
  function closeFactorMantissa (  ) external view returns ( uint256 );
  function comptrollerImplementation (  ) external view returns ( address );
  function distributeVAIMinterVenus ( address vaiMinter ) external;
  function enterMarkets ( address[] memory vTokens ) external returns ( uint256[] memory);
  function exitMarket ( address vTokenAddress ) external returns ( uint256 );
  function getAccountLiquidity ( address account ) external view returns ( uint256, uint256, uint256 );
  function getAllMarkets (  ) external view returns ( address[] memory);
  function getAssetsIn ( address account ) external view returns ( address[] memory);
  function getBlockNumber (  ) external view returns ( uint256 );
  function getHypotheticalAccountLiquidity ( address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount ) external view returns ( uint256, uint256, uint256 );
  function getXVSAddress (  ) external view returns ( address );
  function isComptroller (  ) external view returns ( bool );
  function lastContributorBlock ( address ) external view returns ( uint256 );
  function liquidateBorrowAllowed ( address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount ) external returns ( uint256 );
  function liquidateBorrowVerify ( address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens ) external;
  function liquidateCalculateSeizeTokens ( address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount ) external view returns ( uint256, uint256 );
  function liquidateVAICalculateSeizeTokens ( address vTokenCollateral, uint256 actualRepayAmount ) external view returns ( uint256, uint256 );
  function liquidationIncentiveMantissa (  ) external view returns ( uint256 );
  function markets ( address ) external view returns ( bool isListed, uint256 collateralFactorMantissa, bool isVenus );
  function maxAssets (  ) external view returns ( uint256 );
  function minReleaseAmount (  ) external view returns ( uint256 );
  function mintAllowed ( address vToken, address minter, uint256 mintAmount ) external returns ( uint256 );
  function mintGuardianPaused ( address ) external view returns ( bool );
  function mintVAIGuardianPaused (  ) external view returns ( bool );
  function mintVerify ( address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens ) external;
  function mintedVAIs ( address ) external view returns ( uint256 );
  function oracle (  ) external view returns ( address );
  function pauseGuardian (  ) external view returns ( address );
  function pendingAdmin (  ) external view returns ( address );
  function pendingComptrollerImplementation (  ) external view returns ( address );
  function protocolPaused (  ) external view returns ( bool );
  function redeemAllowed ( address vToken, address redeemer, uint256 redeemTokens ) external returns ( uint256 );
  function redeemVerify ( address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens ) external;
  function releaseStartBlock (  ) external view returns ( uint256 );
  function releaseToVault (  ) external;
  function repayBorrowAllowed ( address vToken, address payer, address borrower, uint256 repayAmount ) external returns ( uint256 );
  function repayBorrowVerify ( address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex ) external;
  function repayVAIGuardianPaused (  ) external view returns ( bool );
  function seizeAllowed ( address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens ) external returns ( uint256 );
  function seizeGuardianPaused (  ) external view returns ( bool );
  function seizeVerify ( address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens ) external;
  function setMintedVAIOf ( address owner, uint256 amount ) external returns ( uint256 );
  function transferAllowed ( address vToken, address src, address dst, uint256 transferTokens ) external returns ( uint256 );
  function transferGuardianPaused (  ) external view returns ( bool );
  function transferVerify ( address vToken, address src, address dst, uint256 transferTokens ) external;
  function treasuryAddress (  ) external view returns ( address );
  function treasuryGuardian (  ) external view returns ( address );
  function treasuryPercent (  ) external view returns ( uint256 );
  function vaiController (  ) external view returns ( address );
  function vaiMintRate (  ) external view returns ( uint256 );
  function vaiVaultAddress (  ) external view returns ( address );
  function venusAccrued ( address ) external view returns ( uint256 );
  function venusBorrowState ( address ) external view returns ( uint224 index, uint32 block );
  function venusBorrowerIndex ( address, address ) external view returns ( uint256 );
  function venusContributorSpeeds ( address ) external view returns ( uint256 );
  function venusInitialIndex (  ) external view returns ( uint224 );
  function venusRate (  ) external view returns ( uint256 );
  function venusSpeeds ( address ) external view returns ( uint256 );
  function venusSupplierIndex ( address, address ) external view returns ( uint256 );
  function venusSupplyState ( address ) external view returns ( uint224 index, uint32 block );
  function venusVAIRate (  ) external view returns ( uint256 );
  function venusVAIVaultRate (  ) external view returns ( uint256 );
}


interface IVBNB {
  function _acceptAdmin (  ) external returns ( uint256 );
  function _reduceReserves ( uint256 reduceAmount ) external returns ( uint256 );
  function _setComptroller ( address newComptroller ) external returns ( uint256 );
  function _setInterestRateModel ( address newInterestRateModel ) external returns ( uint256 );
  function _setPendingAdmin ( address newPendingAdmin ) external returns ( uint256 );
  function _setReserveFactor ( uint256 newReserveFactorMantissa ) external returns ( uint256 );
  function accrualBlockNumber (  ) external view returns ( uint256 );
  function accrueInterest (  ) external returns ( uint256 );
  function admin (  ) external view returns ( address );
  function allowance ( address owner, address spender ) external view returns ( uint256 );
  function approve ( address spender, uint256 amount ) external returns ( bool );
  function balanceOf ( address owner ) external view returns ( uint256 );
  function balanceOfUnderlying ( address owner ) external returns ( uint256 );
  function borrow ( uint256 borrowAmount ) external returns ( uint256 );
  function borrowBalanceCurrent ( address account ) external returns ( uint256 );
  function borrowBalanceStored ( address account ) external view returns ( uint256 );
  function borrowIndex (  ) external view returns ( uint256 );
  function borrowRatePerBlock (  ) external view returns ( uint256 );
  function comptroller (  ) external view returns ( address );
  function decimals (  ) external view returns ( uint8 );
  function exchangeRateCurrent (  ) external returns ( uint256 );
  function exchangeRateStored (  ) external view returns ( uint256 );
  function getAccountSnapshot ( address account ) external view returns ( uint256, uint256, uint256, uint256 );
  function getCash (  ) external view returns ( uint256 );
  function initialize ( address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string memory name_, string memory symbol_, uint8 decimals_ ) external;
  function interestRateModel (  ) external view returns ( address );
  function isVToken (  ) external view returns ( bool );
  function liquidateBorrow ( address borrower, address vTokenCollateral ) external payable;
  function mint (  ) external payable;
  function name (  ) external view returns ( string memory );
  function pendingAdmin (  ) external view returns ( address );
  function redeem ( uint256 redeemTokens ) external returns ( uint256 );
  function redeemUnderlying ( uint256 redeemAmount ) external returns ( uint256 );
  function repayBorrow (  ) external payable;
  function repayBorrowBehalf ( address borrower ) external payable;
  function reserveFactorMantissa (  ) external view returns ( uint256 );
  function seize ( address liquidator, address borrower, uint256 seizeTokens ) external returns ( uint256 );
  function supplyRatePerBlock (  ) external view returns ( uint256 );
  function symbol (  ) external view returns ( string memory );
  function totalBorrows (  ) external view returns ( uint256 );
  function totalBorrowsCurrent (  ) external returns ( uint256 );
  function totalReserves (  ) external view returns ( uint256 );
  function totalSupply (  ) external view returns ( uint256 );
  function transfer ( address dst, uint256 amount ) external returns ( bool );
  function transferFrom ( address src, address dst, uint256 amount ) external returns ( bool );
}

interface IPancakeRouter01 {
    function factory() external pure returns (address);
    function WETH() external pure returns (address);

    function addLiquidity(
        address tokenA,
        address tokenB,
        uint amountADesired,
        uint amountBDesired,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB, uint liquidity);
    function addLiquidityETH(
        address token,
        uint amountTokenDesired,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external payable returns (uint amountToken, uint amountETH, uint liquidity);
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external returns (uint amountA, uint amountB);
    function removeLiquidityETH(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountToken, uint amountETH);
    function removeLiquidityWithPermit(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountA, uint amountB);
    function removeLiquidityETHWithPermit(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountToken, uint amountETH);
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
    function swapTokensForExactTokens(
        uint amountOut,
        uint amountInMax,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)
    external
    payable
    returns (uint[] memory amounts);
    function swapTokensForExactETH(uint amountOut, uint amountInMax, address[] calldata path, address to, uint deadline)
    external
    returns (uint[] memory amounts);
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
    external
    returns (uint[] memory amounts);
    function swapETHForExactTokens(uint amountOut, address[] calldata path, address to, uint deadline)
    external
    payable
    returns (uint[] memory amounts);

    function quote(uint amountA, uint reserveA, uint reserveB) external pure returns (uint amountB);
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external pure returns (uint amountOut);
    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) external pure returns (uint amountIn);
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
    function getAmountsIn(uint amountOut, address[] calldata path) external view returns (uint[] memory amounts);
}

interface IPancakeRouter02 is IPancakeRouter01 {
    function removeLiquidityETHSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline
    ) external returns (uint amountETH);
    function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(
        address token,
        uint liquidity,
        uint amountTokenMin,
        uint amountETHMin,
        address to,
        uint deadline,
        bool approveMax, uint8 v, bytes32 r, bytes32 s
    ) external returns (uint amountETH);

    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external payable;
    function swapExactTokensForETHSupportingFeeOnTransferTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external;
}

interface IVAI {
    function liquidateVAI(address borrower, uint repayAmount, VTokenInterface vTokenCollateral) external returns (uint, uint);
}
