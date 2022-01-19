define(
    "Message",
    [],
    function () {

        function Message(model) {
            if (model !== undefined) { // server response
                this.author = ko.observable(model.author);
                this.body = ko.observable(JSON.stringify(model.body));
                this.file_name = ko.observable(model.file_name);
                this.file_size = ko.observable(model.file_size);
                this.file_type = ko.observable(model.file_type);
                this.file_data = ko.observable(model.file_data);
            } else {
                this.author = ko.observable("Anonymous");
                this.body = ko.observable("{\"id\": \"9352114b-f1dc-4e46-850f-758be49ddb3e\",\"created_at\": null,\"updated_at\": null,\"deleted_at\": null,\"user_name\": \"Vasya\",\"full_name\": \"Pupkin\",\"organization\": \"Home\",\"phone\": \"8 111 222 33 44\",\"email\": \"milo@milo.ru\",\"is_online\": false,\"delegated_until\": null, \"is_block\": false,\"is_ldap\": false,\"federal_districts_id\": \"00000000-0000-0000-0000-000000000000\",\"delegate_from\": \"0001-01-01T00:00:00Z\",\"delegate_to\": \"0001-01-01T00:00:00Z\",\"post\": null}");
                this.file_name = ko.observable("");
                this.file_size = ko.observable(0);
                this.file_type = ko.observable("");
                this.file_data = ko.observable("");
            }

            this.toModel = function () {
                return {
                    author: this.author(),
                    body: this.body(),
                    file_name: this.file_name(),
                    file_size: this.file_size(),
                    file_type: this.file_type(),
                    file_data: this.file_data(),
                };
            }

            this.readFile = async function (file) {
                if (!file) {
                    return new Message();
                }

                const content = await new Promise((resolve, reject) => {
                    var reader = new FileReader();
                    // reader.onload = event => resolve(event.target.result)
                    reader.onload = event => resolve(event.target.result.replace(/^.*base64,/, ''))
                    // reader.readAsText(file);
                    reader.readAsDataURL(file);
                });

                this.file_name(file.name);
                this.file_size(file.size);
                this.file_type(file.type);
                this.file_data(content);
                return this;
            }
        }

        return Message;
    }
);
