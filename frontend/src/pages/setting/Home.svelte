<script>
    import Panel from "../../components/Panel.svelte";

    export let token = ""
    export let master = ""
    export let maintenance_start_field = ""
    export let maintenance_end_field = ""
    let start_maintenance = ""
    let end_maintenance = ""
    let css_loader = "display: none;";
    let msgloader = "";
    async function saveEntry() {
        let flag = false;
        let msg = "";
        if (start_maintenance == "") {
            flag = true;
            msg += "The Jam Maintenance is required\n";
        }
        if (end_maintenance == "") {
            flag = true;
            msg += "The Jam Maintenance is required\n";
        }
        if (flag == false) {
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/savesetting", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    master: master,
                    maintenance_start: start_maintenance,
                    maintenance_end: end_maintenance,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    const handleKeyboard_time = () => {
        let numbera;
        for (let i = 0; i < start_maintenance.length; i++) {
            numbera = parseInt(start_maintenance[i]);
            if (isNaN(numbera)) {
                if (start_maintenance[i] != ":") {
                    start_maintenance = "";
                }
            }
        }
        for (let i = 0; i < end_maintenance.length; i++) {
            numbera = parseInt(end_maintenance[i]);
            if (isNaN(numbera)) {
                if (end_maintenance[i] != ":") {
                    end_maintenance = "";
                }
            }
        }
    };
    $:{
        start_maintenance = maintenance_start_field
        end_maintenance = maintenance_end_field
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-md-4">
            <Panel
                card_footer={false}
                css_body="" 
                height_body="250px">
                <slot:template slot="cheader">
                    Maintenance
                    <div class="float-end">
                        <button
                            on:click={() => {
                                saveEntry();
                            }}
                            class="btn btn-warning btn-sm"
                            style="border-radius: 0px;"
                        >
                            Save
                        </button>
                    </div>
                </slot:template>
                <slot:template slot="cbody">
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label"
                            >Start</label
                        >
                        <input
                            bind:value={start_maintenance}
                            on:keyup={handleKeyboard_time}
                            minlength="0"
                            maxlength="8"
                            type="text"
                            style="text-align:center;"
                            class="form-control required"
                            placeholder="Start Maintenance"
                            aria-label="Start Maintenance"
                        />
                    </div>
                    <div class="mb-3">
                        <label
                            for="exampleFormControlInput1"
                            class="form-label">End</label
                        >
                        <input
                            bind:value={end_maintenance}
                            on:keyup={handleKeyboard_time}
                            minlength="0"
                            maxlength="8"
                            type="text"
                            style="text-align:center;"
                            class="form-control required"
                            placeholder="End Maintenance"
                            aria-label="End Maintenance"
                        />
                    </div>
                </slot:template>
                <slot:template slot="cfooter">
                    
                </slot:template>
            </Panel>
        </div>
    </div>
</div>