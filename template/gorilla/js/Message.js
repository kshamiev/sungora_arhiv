define(
    "Message",
    [],
    function () {

        function Message(model) {
            if (model !== undefined) { // server response
                this.author = ko.observable(model.author);
                this.body = ko.observable(JSON.stringify(model.body));
            } else {
                this.author = ko.observable("Anonymous");
                this.body = ko.observable("{\"id\": \"9352114b-f1dc-4e46-850f-758be49ddb3e\",\"created_at\": null,\"updated_at\": null,\"deleted_at\": null,\"user_name\": \"Vasya\",\"full_name\": \"Pupkin\",\"organization\": \"Home\",\"phone\": \"8 111 222 33 44\",\"email\": \"milo@milo.ru\",\"is_online\": false,\"delegated_until\": null, \"is_block\": false,\"is_ldap\": false,\"federal_districts_id\": \"00000000-0000-0000-0000-000000000000\",\"delegate_from\": \"0001-01-01T00:00:00Z\",\"delegate_to\": \"0001-01-01T00:00:00Z\",\"post\": null}");
            }

            this.toModel = function () {
                return {
                    author: this.author(),
                    body: this.body()
                };
            }
        }

        return Message;
    }
);
