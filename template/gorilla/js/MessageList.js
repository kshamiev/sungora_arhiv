define(
    "MessageList",
    [
        "Message"
    ],
    function (Message) {

        function MessageList(ws) {
            var that = this;
            this.messages = ko.observableArray();
            this.editingMessage = ko.observable(new Message());
            this.send = function () {
                var model = this.editingMessage().toModel();
                model.body = JSON.parse(model.body);
                ws.send($.toJSON(model));
                var message = new Message();
                this.editingMessage(message);
            };
            ws.onmessage = function (e) {
                var model = $.evalJSON(e.data);
                var msg = new Message(model);
                that.messages.push(msg);
                console.log(model)
            };
        }

        return MessageList;
    }
);
