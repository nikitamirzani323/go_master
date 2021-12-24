<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
    export let token = ""
    export let master = ""
    export let listInvoice = []
    export let totalrecord = 0

    let dispatch = createEventDispatcher();
    let page = "Invoice";
    let screen_height = screen.height - 480;
    let sData = "New";
    let myModal_newdata = "";
    let select_periode = "";
    let searchInvoice = "";
    let filterinvoice = [];
    let css_loader = "display: none;";
    let msgloader = "";
    $: {
        if (searchInvoice) {
            filterinvoice = listInvoice.filter((item) =>
                item.invoice_id
                    .toLowerCase()
                    .includes(searchInvoice.toLowerCase()) || 
                item.invoice_company
                    .toLowerCase()
                    .includes(searchInvoice.toLowerCase())
            );
        } else {
            filterinvoice = [...listInvoice];
        }
    }
    async function saveEntry() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (select_periode == "") {
            flag = true;
            msg += "The Periode is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/saveinvoice", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    tipe: "",
                    master: master,
                    periode: select_periode,
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
            RefreshHalaman();
            select_periode = "";
            myModal_newdata.hide();
        } else {
            alert(msg);
        }
    }
    async function UpdateWinlose(e) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/saveinvoicewinlose", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                tipe: "UPDATE-WINLOSE",
                invoice: e,
                master: master,
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
        RefreshHalaman();
    }
    async function UpdateStatus(e) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/saveinvoicewinlose", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                tipe: "UPDATE-STATUS",
                invoice: e,
                master: master,
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
        RefreshHalaman();
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        myModal_newdata = new bootstrap.Modal(document.getElementById("modalNew"));
        myModal_newdata.show();
    };
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-md-12">
            <button
                on:click={() => {
                    RefreshHalaman();
                }}
                class="btn btn-primary btn-sm"
                style="border-radius: 0px;">
                Refresh
            </button>
            <button
                on:click={() => {
                    NewData();
                }}
                class="btn btn-primary btn-sm"
                style="border-radius: 0px;">
                New
            </button>
            <Panel height_body="{screen_height}px">
                <slot:template slot="cheader">
                    {page}
                </slot:template>
                <slot:template slot="csearch">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchInvoice}
                            type="text"
                            class="form-control"
                            placeholder="Search Invoice"
                            aria-label="Search"
                        />
                    </div>
                </slot:template>
                <slot:template slot="cbody">
                    {#if totalrecord > 0}
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th
                                        colspan="2"
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">&nbsp;</th>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">NO</th>
                                    <th
                                        NOWRAP
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">&nbsp;</th>
                                    <th
                                        NOWRAP
                                        width="5%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">DATE</th>
                                    <th
                                        NOWRAP
                                        width="5%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">INVOICE</th>
                                    <th
                                        NOWRAP
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">COMPANY</th>
                                    <th
                                        NOWRAP
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">PERIODE</th>
                                    <th
                                        NOWRAP
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">WINLOSE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterinvoice as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            title="Update winlose"
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;">
                                            {#if rec.invoice_status == "PROGRESS"}
                                            <i
                                                on:click={() => {
                                                    UpdateWinlose(rec.invoice_id);
                                                }}
                                                class="bi bi-arrow-repeat"></i>
                                            {/if}
                                        </td>
                                        <td
                                            NOWRAP
                                            title="Update Status"
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;">
                                            {#if rec.invoice_status == "PROGRESS"}
                                            <i 
                                                on:click={() => {
                                                    UpdateStatus(rec.invoice_id);
                                                }}
                                                class="bi bi-save-fill"></i>
                                            {/if}
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;">{rec.invoice_no}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.invoice_statuscss}">{rec.invoice_status}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;">{rec.invoice_date}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.invoice_id}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.invoice_company}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.invoice_name}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.invoice_winlose_css}">
                                                {new Intl.NumberFormat().format(
                                                    rec.invoice_winlose
                                                )}</td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    {:else}
                        <center>
                            <Loader />
                        </center>
                    {/if}
                </slot:template>
                <slot:template slot="cfooter">
                    Total Record : {totalrecord}
                </slot:template>
            </Panel>
        </div>
    </div>
</div>
<Modal
    modal_id={"modalNew"}
    modal_size={"modal-dialog-centered"}
    modal_body_height={"height:300px;"}
    modal_footer_flag={true}
>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">Entry/{sData}</h5>
    </slot:template>
    <slot:template slot="body">
        <div class="container">
            <div class="row">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Periode</label>
                    <select 
                        bind:value="{select_periode}"
                        class="form-control">
                        <option value="01">JANUARY</option>
                        <option value="02">FEBUARY</option>
                        <option value="03">MARET</option>
                        <option value="04">APRIL</option>
                        <option value="05">MAY</option>
                        <option value="06">JUNE</option>
                        <option value="07">JULY</option>
                        <option value="08">AUGUSTUS</option>
                        <option value="09">SEPTEMBER</option>
                        <option value="10">OCTOBER</option>
                        <option value="11">NOVEMBER</option>
                        <option value="12">DECEMBER</option>
                    </select>
                </div>
                
            </div>
        </div>
    </slot:template>
    <slot:template slot="footer">
        <div class="float-end">
            <button
                on:click={() => {
                    saveEntry();
                }}
                class="btn btn-warning"
                style="border-radius: 0px;"
            >
                Save
            </button>
        </div>
    </slot:template>
</Modal>