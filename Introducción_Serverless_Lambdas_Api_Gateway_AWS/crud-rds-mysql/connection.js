const configDB = {
    host: 'localhost',
    user: 'admin',
    password: '',
    port: '3306',
    database: 'curso_sls',
    debug: true
}

function initializeConnection(config) {
    function addDisconnectHandler(connection) {
        connection.on("error", function (error) {
            if (error instanceof Error) {
                if (error.code === "PROTOCOL_CONNECTION_LOST") {
                    console.log(error.stack);
                    console.log("Lost connection. Reconnecting...");
                    initializeConnection(connection.config);
                } else if (error.fatal) {
                    throw error;
                }
            }
        });
    }
    const connection = mysql.createConnection(config);

    addDisconnectHandler(connection);

    connection.connect();

    return connection;
}

const connection = initializeConnection(configDB);

module.exports = connection;
