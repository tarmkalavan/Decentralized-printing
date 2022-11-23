const printerJsonAbi = `[{"inputs":[{"internalType":"string","name":"_displayName","type":"string"},{"internalType":"string","name":"_printerName","type":"string"},{"internalType":"uint256","name":"_price","type":"uint256"},{"internalType":"string","name":"_location","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"inputs":[],"name":"acceptError","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newTx","type":"address"}],"name":"addToQueue","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"clearance","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"dismissError","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"finished","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getFrontQueue","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getPrice","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getPrinterState","outputs":[{"internalType":"enum PrinterState","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getQueue","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"notifyError","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"printerData","outputs":[{"internalType":"string","name":"displayName","type":"string"},{"internalType":"string","name":"printerName","type":"string"},{"internalType":"uint256","name":"price","type":"uint256"},{"internalType":"string","name":"location","type":"string"},{"internalType":"address","name":"onGoing","type":"address"},{"internalType":"enum PrinterState","name":"state","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"enum PrinterState","name":"state","type":"uint8"}],"name":"updatePrinterState","outputs":[],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]`;
const transactionJsonAbi = `[{"inputs":[{"internalType":"string","name":"_linkFile","type":"string"},{"internalType":"address","name":"_printer","type":"address"},{"internalType":"uint256","name":"_lenPage","type":"uint256"}],"stateMutability":"payable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"inputs":[],"name":"clearance","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getTxState","outputs":[{"internalType":"enum TxState","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"refund","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"reportError","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"transactionData","outputs":[{"internalType":"string","name":"linkFile","type":"string"},{"internalType":"uint256","name":"price","type":"uint256"},{"internalType":"enum TxState","name":"state","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"enum TxState","name":"state","type":"uint8"}],"name":"updateTxState","outputs":[],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]`;
const centralServerJsonAbi = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getPrinters","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"printerArr","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newPrinter","type":"address"}],"name":"registerPrinter","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"printerAddr","type":"address"}],"name":"removePrinter","outputs":[],"stateMutability":"nonpayable","type":"function"}]`;
const Abi = {
    printer: printerJsonAbi,
    transaction: transactionJsonAbi,
    centralServer: centralServerJsonAbi,
};

export default Abi;
