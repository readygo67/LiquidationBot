
// File: @openzeppelin/contracts/utils/Address.sol


// OpenZeppelin Contracts (last updated v4.5.0) (utils/Address.sol)

pragma solidity ^0.8.1;

/**
 * @dev Collection of functions related to the address type
 */
library Address {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://diligence.consensys.net/posts/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.5.11/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCall(target, data, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        require(isContract(target), "Address: call to non-contract");

        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        require(isContract(target), "Address: static call to non-contract");

        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionDelegateCall(target, data, "Address: low-level delegate call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(isContract(target), "Address: delegate call to non-contract");

        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verifies that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            // Look for revert reason and bubble it up if present
            if (returndata.length > 0) {
                // The easiest way to bubble the revert reason is using memory via assembly

                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}

// File: @openzeppelin/contracts/utils/Context.sol


// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

pragma solidity ^0.8.0;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

// File: @openzeppelin/contracts/access/Ownable.sol


// OpenZeppelin Contracts v4.4.1 (access/Ownable.sol)

pragma solidity ^0.8.0;


/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}

// File: LiquidationBot/contracts/SafeMath.sol


// OpenZeppelin Contracts v4.4.1 (utils/math/SafeMath.sol)

pragma solidity ^0.8.0;

// CAUTION
// This version of SafeMath should only be used with Solidity 0.8 or later,
// because it relies on the compiler's built in overflow checks.

/**
 * @dev Wrappers over Solidity's arithmetic operations.
 *
 * NOTE: `SafeMath` is generally not needed starting with Solidity 0.8, since the compiler
 * now has built in overflow checking.
 */
library SafeMath {
    /**
     * @dev Returns the addition of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryAdd(uint256 a, uint256 b) internal pure returns (bool, uint256) {
    unchecked {
        uint256 c = a + b;
        if (c < a) return (false, 0);
        return (true, c);
    }
    }

    /**
     * @dev Returns the substraction of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function trySub(uint256 a, uint256 b) internal pure returns (bool, uint256) {
    unchecked {
        if (b > a) return (false, 0);
        return (true, a - b);
    }
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryMul(uint256 a, uint256 b) internal pure returns (bool, uint256) {
    unchecked {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
        if (a == 0) return (true, 0);
        uint256 c = a * b;
        if (c / a != b) return (false, 0);
        return (true, c);
    }
    }

    /**
     * @dev Returns the division of two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryDiv(uint256 a, uint256 b) internal pure returns (bool, uint256) {
    unchecked {
        if (b == 0) return (false, 0);
        return (true, a / b);
    }
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryMod(uint256 a, uint256 b) internal pure returns (bool, uint256) {
    unchecked {
        if (b == 0) return (false, 0);
        return (true, a % b);
    }
    }

    /**
     * @dev Returns the addition of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `+` operator.
     *
     * Requirements:
     *
     * - Addition cannot overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        return a + b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return a - b;
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `*` operator.
     *
     * Requirements:
     *
     * - Multiplication cannot overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        return a * b;
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator.
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return a / b;
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting when dividing by zero.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        return a % b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting with custom message on
     * overflow (when the result is negative).
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {trySub}.
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
    unchecked {
        require(b <= a, errorMessage);
        return a - b;
    }
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting with custom message on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
    unchecked {
        require(b > 0, errorMessage);
        return a / b;
    }
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting with custom message when dividing by zero.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {tryMod}.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(
        uint256 a,
        uint256 b,
        string memory errorMessage
    ) internal pure returns (uint256) {
    unchecked {
        require(b > 0, errorMessage);
        return a % b;
    }
    }
}
// File: LiquidationBot/contracts/interface.sol


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
    function borrowBalanceStored(address account) external view returns (uint);
    function exchangeRateCurrent() external returns (uint);
    function exchangeRateStored() external view returns (uint);
    function getCash() external view returns (uint);
    function accrueInterest() external returns (uint);
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

// File: LiquidationBot/contracts/PancakeLibrary.sol


pragma solidity ^0.8;



library PancakeLibrary {
    using SafeMath for uint;

    // returns sorted token addresses, used to handle return values from pairs sorted in this order
    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        require(tokenA != tokenB, 'PancakeLibrary: IDENTICAL_ADDRESSES');
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        require(token0 != address(0), 'PancakeLibrary: ZERO_ADDRESS');
    }


    // calculates the CREATE2 address for a pair without making any external calls
    function pairFor(address factory, address tokenA, address tokenB) internal pure returns (address pair) {
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        pair = address(uint160(uint(keccak256(abi.encodePacked(
                hex'ff',
                factory,
                keccak256(abi.encodePacked(token0, token1)),
                hex'00fb7f630766e6a796048ea87d01acd3068e8ff67d078148a3fa3f4a84f69bd5' // init code hash
            )))));
    }

    // fetches and sorts the reserves for a pair
    function getReserves(address factory, address tokenA, address tokenB) internal view returns (uint reserveA, uint reserveB) {
        (address token0,) = sortTokens(tokenA, tokenB);
        pairFor(factory, tokenA, tokenB);
        (uint reserve0, uint reserve1,) = IPancakePair(pairFor(factory, tokenA, tokenB)).getReserves();
        (reserveA, reserveB) = tokenA == token0 ? (reserve0, reserve1) : (reserve1, reserve0);
    }

    // given some amount of an asset and pair reserves, returns an equivalent amount of the other asset
    function quote(uint amountA, uint reserveA, uint reserveB) internal pure returns (uint amountB) {
        require(amountA > 0, 'PancakeLibrary: INSUFFICIENT_AMOUNT');
        require(reserveA > 0 && reserveB > 0, 'PancakeLibrary: INSUFFICIENT_LIQUIDITY');
        amountB = amountA.mul(reserveB) / reserveA;
    }

    // given an input amount of an asset and pair reserves, returns the maximum output amount of the other asset
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) internal pure returns (uint amountOut) {
        require(amountIn > 0, 'PancakeLibrary: INSUFFICIENT_INPUT_AMOUNT');
        require(reserveIn > 0 && reserveOut > 0, 'PancakeLibrary: INSUFFICIENT_LIQUIDITY');
        uint amountInWithFee = amountIn.mul(9975);
        uint numerator = amountInWithFee.mul(reserveOut);
        uint denominator = reserveIn.mul(10000).add(amountInWithFee);
        amountOut = numerator / denominator;
    }

    // given an output amount of an asset and pair reserves, returns a required input amount of the other asset
    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) internal pure returns (uint amountIn) {
        require(amountOut > 0, 'PancakeLibrary: INSUFFICIENT_OUTPUT_AMOUNT');
        require(reserveIn > 0 && reserveOut > 0, 'PancakeLibrary: INSUFFICIENT_LIQUIDITY');
        uint numerator = reserveIn.mul(amountOut).mul(10000);
        uint denominator = reserveOut.sub(amountOut).mul(9975);
        amountIn = (numerator / denominator).add(1);
    }

    // performs chained getAmountOut calculations on any number of pairs
    function getAmountsOut(address factory, uint amountIn, address[] memory path) internal view returns (uint[] memory amounts) {
        require(path.length >= 2, 'PancakeLibrary: INVALID_PATH');
        amounts = new uint[](path.length);
        amounts[0] = amountIn;
        for (uint i; i < path.length - 1; i++) {
            (uint reserveIn, uint reserveOut) = getReserves(factory, path[i], path[i + 1]);
            amounts[i + 1] = getAmountOut(amounts[i], reserveIn, reserveOut);
        }
    }

    // performs chained getAmountIn calculations on any number of pairs
    function getAmountsIn(address factory, uint amountOut, address[] memory path) internal view returns (uint[] memory amounts) {
        require(path.length >= 2, 'PancakeLibrary: INVALID_PATH');
        amounts = new uint[](path.length);
        amounts[amounts.length - 1] = amountOut;
        for (uint i = path.length - 1; i > 0; i--) {
            (uint reserveIn, uint reserveOut) = getReserves(factory, path[i - 1], path[i]);
            amounts[i - 1] = getAmountIn(amounts[i], reserveIn, reserveOut);
        }
    }
}
// File: LiquidationBot/contracts/liquidator.sol


pragma solidity ^0.8;





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
    case1
       0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE
       ["0xfD5840Cd36d94D7229439859C0112a4185BC0255","0xfD5840Cd36d94D7229439859C0112a4185BC0255","0x55d398326f99059fF775485246999027B3197955","0x55d398326f99059fF775485246999027B3197955","0x564EE8bF0bA977A1ccc92fe3D683AbF4569c9f5E"]
       21034205123917372652

    case2.1(TODO)




    case2.2
        height15379782, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, repaySmbol:vETH, flashLoanFrom:0x74E4716E431f45807DCF19f284c7aA99F18a4fbc, repayAddress:0xf508fCD89b8bd15579dc79A6827cB4686A3592c8, repayValue:27760775212595059202.7, repayAmount:9834830202498401 seizedSymbol:vETH, seizedAddress:0xf508fCD89b8bd15579dc79A6827cB4686A3592c8, seizedCTokenAmount:53542165, seizedUnderlyingTokenAmount:10818313118937201.583675475128906, seizedUnderlyingTokenValue:30536852440824038910.2407636463629662
        calculateSeizedTokenAmount case2: seizedSymbol == repaySymbol and symbol is not stable coin, account:0xFAbE4C180b6eDad32eA0Cf56587c54417189e422, symbol:vETH, seizedAmount:10818313118937201.583675475128906, returnAmout:9859478899747771, usdtAmount:2702645379426654958, gasFee:2425666800000000000, profit:0.2786001639516656
        case2, profitable liquidation catched:&{0xFAbE4C180b6eDad32eA0Cf56587c54417189e422 0.9512268623785098 15379543 0001-01-01 00:00:00 +0000 UTC}, profit:0.2786001639516656

        0x74E4716E431f45807DCF19f284c7aA99F18a4fbc
        path1: ["0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x55d398326f99059fF775485246999027B3197955"]
        path2: ["0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x55d398326f99059fF775485246999027B3197955"]
        tokens:["0xf508fCD89b8bd15579dc79A6827cB4686A3592c8","0xf508fCD89b8bd15579dc79A6827cB4686A3592c8","0x2170ed0880ac9a755fd29b2688956bd959f933f8","0x2170ed0880ac9a755fd29b2688956bd959f933f8","0xFAbE4C180b6eDad32eA0Cf56587c54417189e422"]
        9834830202498401

    case3.1(TODO)




    case3.2
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


    case4.2 gas:1666282 => 1181926
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

    case5.2 gas:927990 => 818503
        calculateSeizedTokenAmount case5: seizedSymbol and repaySymbol are not stable coin
        height:15415559, account:0x5CA18142b3Aa0E4e580ec939fe3761af37395262, repaySymbol:vTUSD, repayUnderlyingAmount:5012979382087345840, seizedSymbol:vBNB, seizedVTokenAmount:67209325, seizedUnderlyingAmount:14483051888262770.9963578757692161, seizedValue:5504860657675716182.0207406384930878, flashLoanReturnAmout:5025543240187815374, remain:1606269562758565, gasFee:2850673687500000000, profit:-2.2406430009646194
        flashLoanFrom:0x2E28b9B74D6d99D4697e913b82B41ef1CAC51c6C, path1:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x14016E85a25aeb13065688cAFB43044C2ef86784], path2:[0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x55d398326f99059fF775485246999027B3197955], addresses:[0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3 0xA07c5b74C9B40447a954e1466938b865b6BBea36 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c 0x14016E85a25aeb13065688cAFB43044C2ef86784 0x5CA18142b3Aa0E4e580ec939fe3761af37395262]

        flashLoanFrom:0x2E28b9B74D6d99D4697e913b82B41ef1CAC51c6C
        path1:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x14016E85a25aeb13065688cAFB43044C2ef86784"]
        path2:["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x55d398326f99059fF775485246999027B3197955"]
        addresses:["0x08CEB3F4a7ed3500cA0982bcd0FC7816688084c3", "0xA07c5b74C9B40447a954e1466938b865b6BBea36", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x14016E85a25aeb13065688cAFB43044C2ef86784", "0x5CA18142b3Aa0E4e580ec939fe3761af37395262"]

    case5.3 gas:1189122 => 998937
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

    case7.2 gas:2112126 ==> 1371980
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
            (,,uint shortfall) = Comptroller(ComptrollerAddr).getAccountLiquidity(_tokens[4]);
            require(shortfall > 0, "shortfall must greater than zer 0");

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

