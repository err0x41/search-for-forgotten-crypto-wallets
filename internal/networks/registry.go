package networks

type Chain struct {
	Name     string
	Type     string
	RPCs     []string
	Index    int
	Currency string
}

func GetRegistry() []Chain {
	return []Chain{
		// --- EVM SECTOR (30 networks) ---
		{"Ethereum", "evm", []string{"https://rpc.ankr.com/eth"}, 60, "ETH"},
		{"BNB Chain", "evm", []string{"https://rpc.ankr.com/bsc"}, 714, "BNB"},
		{"Polygon", "evm", []string{"https://polygon-rpc.com"}, 60, "MATIC"},
		{"Arbitrum", "evm", []string{"https://arb1.arbitrum.io/rpc"}, 60, "ETH"},
		{"Optimism", "evm", []string{"https://mainnet.optimism.io"}, 60, "ETH"},
		{"Base", "evm", []string{"https://mainnet.base.org"}, 60, "ETH"},
		{"Avalanche", "evm", []string{"https://api.avax.network/ext/bc/C/rpc"}, 60, "AVAX"},
		{"zkSync", "evm", []string{"https://mainnet.era.zksync.io"}, 60, "ETH"},
		{"Linea", "evm", []string{"https://rpc.linea.build"}, 60, "ETH"},
		{"Scroll", "evm", []string{"https://rpc.scroll.io"}, 60, "ETH"},
		{"Blast", "evm", []string{"https://rpc.blast.io"}, 60, "ETH"},
		{"Mantle", "evm", []string{"https://rpc.mantle.xyz"}, 60, "MNT"},
		{"Metis", "evm", []string{"https://andromeda.metis.io/?owner=1088"}, 60, "METIS"},
		{"Gnosis", "evm", []string{"https://rpc.gnosischain.com"}, 60, "xDAI"},
		{"Fantom", "evm", []string{"https://rpc.ftm.tools"}, 60, "FTM"},
		{"Celo", "evm", []string{"https://forno.celo.org"}, 52742, "CELO"},
		{"Cronos", "evm", []string{"https://evm.cronos.org"}, 60, "CRO"},
		{"Kava", "evm", []string{"https://evm.kava.io"}, 60, "KAVA"},
		{"Moonbeam", "evm", []string{"https://rpc.api.moonbeam.network"}, 60, "GLMR"},
		{"Astar", "evm", []string{"https://evm.astar.network"}, 60, "ASTR"},
		{"Harmony", "evm", []string{"https://api.harmony.one"}, 1023, "ONE"},
		{"ETC", "evm", []string{"https://etc.etcdesktop.com"}, 61, "ETC"},
		{"ZetaChain", "evm", []string{"https://zetachain-evm.blockpi.network/v1/rpc/public"}, 60, "ZETA"},
		{"Chiliz", "evm", []string{"https://rpc.chiliz.com"}, 60, "CHZ"},
		{"Ronin", "evm", []string{"https://api.roninchain.com/rpc"}, 60, "RON"},
		{"Rootstock", "evm", []string{"https://public-node.rsk.co"}, 137, "RBTC"},
		{"Manta", "evm", []string{"https://pacific-rpc.manta.network/http"}, 60, "MANTA"},
		{"Mode", "evm", []string{"https://mainnet.mode.network"}, 60, "ETH"},
		{"Klaytn", "evm", []string{"https://public-node-api.klaytnapi.com/v1/cypress"}, 60, "KLAY"},
		{"Fuse", "evm", []string{"https://rpc.fuse.io"}, 60, "FUSE"},

		// --- NON-EVM & OTHERS (20 networks) ---
		{"Bitcoin", "utxo", []string{"https://blockstream.info/api"}, 0, "BTC"},
		{"Litecoin", "utxo", []string{"https://litecoinblockexplorer.net/api"}, 2, "LTC"},
		{"Doge", "utxo", []string{"https://dogechain.info/api"}, 3, "DOGE"},
		{"Bitcoin Cash", "utxo", []string{"https://bch-chain.api.btc.com"}, 145, "BCH"},
		{"Solana", "sol", []string{"https://api.mainnet-beta.solana.com"}, 501, "SOL"},
		{"TON", "ton", []string{"https://toncenter.com/api/v2/jsonRPC"}, 607, "TON"},
		{"Tron", "tron", []string{"https://api.trongrid.io"}, 195, "TRX"},
		{"Near", "sol", []string{"https://rpc.mainnet.near.org"}, 397, "NEAR"},
		{"Sui", "sol", []string{"https://fullnode.mainnet.sui.io"}, 784, "SUI"},
		{"Aptos", "sol", []string{"https://fullnode.mainnet.aptoslabs.com/v1"}, 637, "APT"},
		{"Cardano", "sol", []string{"https://api.koios.rest/api/v0"}, 1815, "ADA"},
		{"Ripple", "sol", []string{"https://s1.ripple.com:51234"}, 144, "XRP"},
		{"Stellar", "sol", []string{"https://horizon.stellar.org"}, 148, "XLM"},
		{"Polkadot", "substrate", []string{"https://rpc.polkadot.io"}, 354, "DOT"},
		{"Cosmos", "sol", []string{"https://rpc.cosmos.network"}, 118, "ATOM"},
		{"Algorand", "sol", []string{"https://mainnet-api.algonode.cloud"}, 283, "ALGO"},
		{"Tezos", "sol", []string{"https://mainnet.api.tez.ie"}, 1729, "XTZ"},
		{"Hedera", "sol", []string{"https://mainnet-public.mirrornode.hedera.com"}, 3030, "HBAR"},
		{"Sei", "sol", []string{"https://sei-rpc.polkachu.com"}, 118, "SEI"},
		{"Celestia", "sol", []string{"https://rpc.lucina-mainnet.celestia-node.com"}, 118, "TIA"},
	}
}
