define(
    "main",
    [
        "MessageList"
    ],
    function (MessageList) {
        var ws = new WebSocket("ws://sungora.ru/api/sun/websocket/gorilla/32be14df-210c-4cb5-8bba-7b76b4cdce98");
        var list = new MessageList(ws);
        ko.applyBindings(list);
    }
);
