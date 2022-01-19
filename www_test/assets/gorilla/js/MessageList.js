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
            this.send = async function () {
                const file = document.getElementById("inputFile").files[0];
                const msg = await this.editingMessage().readFile(file);
                const model = msg.toModel()
                // model.message = JSON.parse(model.message);
                console.log(model)
                ws.send($.toJSON(model));
                const message = new Message();
                this.editingMessage(message);
            };
            ws.onmessage = async function (e) {
                var model = $.evalJSON(e.data);
                var msg = new Message(model);
                that.messages.push(msg);
            };
        }

        return MessageList;
    }
);
