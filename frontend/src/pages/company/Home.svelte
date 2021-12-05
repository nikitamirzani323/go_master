<script>
    import { Icon, Row, Col, Container } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
    export let listCompany = [];
    export let totalrecord = 0;
    export let token = "";
    export let master = "";
    let dispatch = createEventDispatcher();
    let page = "Company";
    let screen_height = screen.height - 480;
    let sData = "New";
    let company_id_field = "";
    let company_name_field = "";
    let company_url_field = "";
    let searchCompany = "";
    let filterCompany = [];
    let css_loader = "display: none;";
    let msgloader = "";
    $: {
        if (searchCompany) {
            filterCompany = listCompany.filter((item) =>
                item.company_name
                    .toLowerCase()
                    .includes(searchCompany.toLowerCase())
            );
        } else {
            filterCompany = [...listCompany];
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e) => {
        const company = {
            e,
        };
        dispatch("handleEditData", company);
    };
    const DetailData = (e) => {
        const company = {
            e,
        };
        dispatch("handleDetailData", company);
    };
    const NewData = () => {
        sData = "New";
        let myModal = new bootstrap.Modal(document.getElementById("modalNew"));
        myModal.show();
    };
    async function saveEntry() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (company_name_field == "") {
            flag = true;
            msg += "The Company is required\n";
        }
        if (company_url_field == "") {
            flag = true;
            msg += "The Url is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savecompany", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master: master,
                    company: company_id_field,
                    name: company_name_field,
                    urldomain: company_url_field,
                    status: "ACTIVE",
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
            company_id_field = "";
            company_name_field = "";
            company_url_field = "";
        } else {
            alert(msg);
        }
    }
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<Container style="margin-top: 70px;">
    <Row>
        <Col>
            <button
                on:click={() => {
                    RefreshHalaman();
                }}
                class="btn btn-primary btn-sm"
                style="border-radius: 0px;"
            >
                Refresh
            </button>
            <button
                on:click={() => {
                    NewData();
                }}
                class="btn btn-primary btn-sm"
                style="border-radius: 0px;"
            >
                New
            </button>
            <Panel height_body="{screen_height}px">
                <slot:template slot="cheader">
                    {page}
                </slot:template>
                <slot:template slot="csearch">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchCompany}
                            type="text"
                            class="form-control"
                            placeholder="Search Company"
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
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >&nbsp;</th
                                    >
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >NO</th
                                    >
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >&nbsp;</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >START</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >END</th
                                    >
                                    <th
                                        width="5%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >ID</th
                                    >
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >COMPANY</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >PERIODE</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >WINLOSE</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >WINLOSE_TEMP</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >SELISIH</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterCompany as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                EditData(rec.company_idcompany);
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;"
                                        >
                                            <Icon name="pencil" />
                                        </td>
                                        <td
                                            on:click={() => {
                                                DetailData(rec.company_idcompany);
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;" >
                                            <i class="bi bi-files"></i>
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.no}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.company_statuscss}"
                                            >{rec.company_status}</td
                                        >

                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.company_startjoin}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.company_endjoin}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.company_idcompany}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.company_name}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.company_periode}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_winlosecss}">
                                                {new Intl.NumberFormat().format(
                                                    rec.company_winlose
                                                )}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_winlosecss}">
                                                {new Intl.NumberFormat().format(
                                                    rec.company_winlosetemp
                                                )}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_selisihwinlosecss}">
                                                {new Intl.NumberFormat().format(
                                                    rec.company_selisihwinlose
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
        </Col>
    </Row>
</Container>
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
        <Container>
            <Row>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">ID</label>
                    <input
                        bind:value={company_id_field}
                        type="text"
                        minlength="3"
                        maxlength="10"
                        class="form-control required"
                        placeholder="ID"
                        aria-label="ID"
                    />
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <input
                        bind:value={company_name_field}
                        type="text"
                        maxlength="70"
                        class="form-control required"
                        placeholder="Name"
                        aria-label="Name"
                    />
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">URL</label>
                    <input
                        bind:value={company_url_field}
                        type="text"
                        class="form-control required"
                        placeholder="URL"
                        aria-label="URL"
                    />
                </div>
            </Row>
        </Container>
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
