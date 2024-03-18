CREATE DATABASE IF NOT EXISTS uniswap_db;
USE uniswap_db;

CREATE TABLE IF NOT EXISTS transactions (
		id VARCHAR(64) NOT NULL,
		blockNumber BIGINT UNSIGNED,
		timeStamp BIGINT UNSIGNED,
		hash CHAR(66),
		nonce INT,
		blockHash CHAR(66),
		fromAddress CHAR(42),
		contractAddress CHAR(42),
		toAddress CHAR(42),
		value VARCHAR(50),
		tokenName VARCHAR(50),
		tokenSymbol VARCHAR(50),
		tokenDecimal INT,
		transactionIndex INT,
		gas INT,
		gasPrice BIGINT,
		gasUsed INT,
		cumulativeGasUsed BIGINT,
		input VARCHAR(50),
		confirmations INT,
		ethPrice DECIMAL(20, 8),
		PRIMARY KEY(id), 
		INDEX (timeStamp),
		INDEX (hash),
		INDEX idx_txn_from_to (hash, fromAddress, toAddress)
);

CREATE TABLE IF NOT EXISTS swaps (
	id VARCHAR(64) NOT NULL,
	txnHash CHAR(66),
	sqrtPriceX96 VARCHAR(100),
	fromAddress CHAR(42),
	toAddress CHAR(42),
	PRIMARY KEY(id), 
	INDEX idx_txn_from_to (txnHash, fromAddress, toAddress)
);