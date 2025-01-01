-- Create the Wallets table
-- This table stores information about wallets.
CREATE TABLE wallets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the wallet
    name VARCHAR(255) NOT NULL,                    -- Wallet name (e.g., user-friendly label)
    spending_pub_key VARCHAR(255) NOT NULL,        -- Spending public key of the wallet
    viewing_pub_key VARCHAR(255) NOT NULL,         -- Viewing public key of the wallet
    viewing_priv_key VARCHAR(255) NOT NULL,        -- Viewing private key of the wallet
    created_at TIMESTAMP DEFAULT NOW(),            -- Timestamp of wallet creation
    updated_at TIMESTAMP DEFAULT NOW()             -- Timestamp of the last update to the wallet
);


-- Create the Addresses table
-- This table stores blockchain addresses associated with wallets.
CREATE TABLE addresses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the address
    wallet_id UUID REFERENCES wallets(id) ON DELETE CASCADE, -- Associated wallet ID
    address VARCHAR(255) NOT NULL UNIQUE,           -- Blockchain address
    chain VARCHAR(50) NOT NULL,                     -- Blockchain name (e.g., eth, bsc)
    type ENUM('token', 'nft') NOT NULL,             -- Type of the asset held by the address
    contract_address VARCHAR(255) NOT NULL,         -- Token or NFT collection contract address
    token_metadata_id UUID REFERENCES token_metadata(id) ON DELETE SET NULL, -- Nullable: Associated token metadata
    collection_metadata_id UUID REFERENCES collection_metadata(id) ON DELETE SET NULL, -- Nullable: Associated NFT collection metadata
    amount BIGINT,                                  -- Amount of tokens or NFTs held
    nft_id VARCHAR(255),                            -- NFT ID for NFT addresses
    created_at TIMESTAMP DEFAULT NOW(),             -- Timestamp of address creation
    updated_at TIMESTAMP DEFAULT NOW()              -- Timestamp of the last update to the address
);

-- Create the TokenMetadata table
-- This table stores general information about tokens, shared across chains.
CREATE TABLE token_metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the token metadata
    token_symbol VARCHAR(50) NOT NULL,             -- Symbol of the token (e.g., USDT, ETH)
    token_name VARCHAR(255) NOT NULL,              -- Full name of the token (e.g., Tether USD, Ethereum)
    coingecko_id VARCHAR(255),                     -- Coingecko ID for querying price data
    logo VARCHAR(255),                             -- URL for the token's logo
    price DECIMAL(20,8),                           -- Current price of the token (NULL if unavailable)
    last_price_updated_at TIMESTAMP,               -- Timestamp when the price was last updated
    created_at TIMESTAMP DEFAULT NOW(),            -- Timestamp of token metadata creation
    updated_at TIMESTAMP DEFAULT NOW()             -- Timestamp of the last update to the metadata
);

-- Create the NFTs table
CREATE TABLE collection_metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the NFT collection
    collection_name VARCHAR(255),                  -- Name of the NFT collection (e.g., CryptoPunks)
    description TEXT,                              -- Description of the NFT collection
    logo_url VARCHAR(255),                         -- URL to the collection's logo
    floor_price DECIMAL(20,8),                     -- Floor price of the NFT (NULL if unavailable)
    metadata_url VARCHAR(255),                     -- URL to the metadata JSON for the collection
    chain VARCHAR(50) NOT NULL,                    -- Blockchain name (e.g., Ethereum, Polygon)
    contract_address VARCHAR(255) NOT NULL,        -- Contract address of the NFT collection
    created_at TIMESTAMP DEFAULT NOW(),            -- Timestamp of record creation
    updated_at TIMESTAMP DEFAULT NOW(),            -- Timestamp of the last update,
    UNIQUE (chain, contract_address)               -- Ensure uniqueness for collections on the same chain
);

-- Create the Transactions table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the transaction
    wallet_id UUID REFERENCES wallets(id) ON DELETE CASCADE, -- Wallet ID of the transaction
    interact_wallet_ID UUID REFERENCES wallets(id) ON DELETE SET NULL, -- Wallet ID of the receiver
    token_metadata_id UUID REFERENCES token_metadata(id) ON DELETE SET NULL, -- Nullable: Associated token metadata
    collection_metadata_id UUID REFERENCES collection_metadata(id) ON DELETE SET NULL, -- Nullable: Associated NFT collection metadata
    nft_id VARCHAR(255),                            -- NFT ID for NFT transactions
    chain VARCHAR(50) NOT NULL,                     -- Blockchain name (e.g., Ethereum, BSC)
    tx_hash VARCHAR(255) NOT NULL UNIQUE,           -- Transaction hash
    type ENUM('send', 'receive', 'interact') NOT NULL, -- Transaction type
    amount DECIMAL(20,8),                           -- Transaction amount (nullable for NFT transactions)
    created_at TIMESTAMP DEFAULT NOW(),             -- Timestamp of the transaction creation
);
