<script>
    import { Icon, Row, Col, Container } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
    export let listPasaran = [];
    export let totalrecord = 0;
    export let token = "";
    export let master = "";
    let dispatch = createEventDispatcher();
    let page = "Pasaran";
    let screen_height = screen.height - 420;
    let sData = "New";
    let idrecord = ""
    let pasaran_nmpasarantogel_field = ""
    let pasaran_tipepasaran_field = ""
    let pasaran_urlpasaran_field = ""
    let pasaran_pasarandiundi_field = ""
    let pasaran_jamtutup_field = ""
    let pasaran_jamjadwal_field = ""
    let pasaran_jamopen_field = ""
    let searchPasaran = "";
    let filterPasaran = [];
    let css_loader = "display: none;";
    let msgloader = "";
    $: {
        if (searchPasaran) {
            filterPasaran = listPasaran.filter((item) =>
                item.pasaran_nmpasarantogel
                    .toLowerCase()
                    .includes(searchPasaran.toLowerCase())
            );
        } else {
            filterPasaran = [...listPasaran];
        }
    }
    const RefreshHalaman = () => {
        idrecord = ""
        pasaran_nmpasarantogel_field = ""
        pasaran_tipepasaran_field = ""
        pasaran_urlpasaran_field = ""
        pasaran_pasarandiundi_field = ""
        pasaran_jamtutup_field = ""
        pasaran_jamjadwal_field = ""
        pasaran_jamopen_field = ""
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e) => {
        const company = {
            e,
        };
        dispatch("handleEditData", company);
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
        if (idrecord == "") {
            flag = true;
            msg += "The ID Pasaran Togel is required\n";
        }
        if (pasaran_nmpasarantogel_field == "") {
            flag = true;
            msg += "The Pasaran Togel is required\n";
        }
        if (pasaran_urlpasaran_field == "") {
            flag = true;
            msg += "The Situs is required\n";
        }
        if (pasaran_pasarandiundi_field == "") {
            flag = true;
            msg += "The Hari Diundi is required\n";
        }
        if (pasaran_jamtutup_field == "") {
            flag = true;
            msg += "The Jam Tutup is required\n";
        }
        if (pasaran_jamjadwal_field == "") {
            flag = true;
            msg += "The Jam Jadwal is required\n";
        }
        if (pasaran_jamopen_field == "") {
            flag = true;
            msg += "The Jam Open is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    master: master,
                    idrecord: idrecord,
                    pasaran_name: pasaran_nmpasarantogel_field,
                    pasaran_diundi: pasaran_pasarandiundi_field,
                    pasaran_url: pasaran_urlpasaran_field,
                    pasaran_jamtutup: pasaran_jamtutup_field,
                    pasaran_jamjadwal: pasaran_jamjadwal_field,
                    pasaran_jamopen: pasaran_jamopen_field,
                    pasaran_tipe: pasaran_tipepasaran_field,
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
            RefreshHalaman()
        } else {
            alert(msg);
        }
    }
    const handleKeyboard_time = (e) => {
        let numbera;
        for (let i = 0; i < pasaran_jamtutup_field.length; i++) {
            numbera = parseInt(pasaran_jamtutup_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_jamtutup_field[i] != ":") {
                    pasaran_jamtutup_field = "";
                }
            }
        }
        for (let i = 0; i < pasaran_jamjadwal_field.length; i++) {
            numbera = parseInt(pasaran_jamjadwal_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_jamjadwal_field[i] != ":") {
                    pasaran_jamjadwal_field = "";
                }
            }
        }
        for (let i = 0; i < pasaran_jamopen_field.length; i++) {
            numbera = parseInt(pasaran_jamopen_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_jamopen_field[i] != ":") {
                    pasaran_jamopen_field = "";
                }
            }
        }
    };
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
                            bind:value={searchPasaran}
                            type="text"
                            class="form-control"
                            placeholder="Search Pasaran"
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
                                        >TYPE</th
                                    >
                                    <th
                                        width="1%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >ID</th
                                    >
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >PASARAN</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >HARI DIUNDI</th
                                    >
                                    <th
                                        width="5%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >TUTUP</th
                                    >
                                    <th
                                        width="5%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >JADWAL</th
                                    >
                                    <th
                                        width="5%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >OPEN</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterPasaran as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                EditData(rec.pasaran_idpasarantogel);
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;"
                                        >
                                            <Icon name="pencil" />
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_no}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_tipepasaran}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_idpasarantogel}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;{rec.company_statuscss}"
                                            >{rec.pasaran_nmpasarantogel}</td
                                        >
                                        
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_pasarandiundi}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_jamtutup}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_jamjadwal}</td
                                        >

                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_jamopen}</td
                                        >
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
    modal_size={"modal-dialog-centered modal-lg"}
    modal_body_height={"height:400px;"}
    modal_footer_flag={true}
>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">Entry/{sData}</h5>
    </slot:template>
    <slot:template slot="body">
        <Container>
            <Row>
                <Col xs="6">
                    <div class="mb-3">
                        <label
                            for="exampleform"
                            class="form-label">ID</label
                        >
                        <input
                            bind:value="{idrecord}"
                            type="text"
                            maxlength="10"
                            class="form-control required"
                            placeholder="ID Pasaran Togel"
                            aria-label="ID Pasaran Togel"
                        />
                    </div>
                    <div class="mb-3">
                        <label
                            for="exampleform"
                            class="form-label">Pasaran</label
                        >
                        <input
                            bind:value="{pasaran_nmpasarantogel_field}"
                            type="text"
                            class="form-control required"
                            placeholder="Pasaran Togel"
                            aria-label="Pasaran Togel"
                        />
                    </div>
                    <div class="mb-3">
                        <label
                            for="exampleform"
                            class="form-label">Situs</label
                        >
                        <input
                            bind:value="{pasaran_urlpasaran_field}"
                            type="text"
                            class="form-control required"
                            placeholder="Situs"
                            aria-label="Situs"
                        />
                    </div>
                    <div class="mb-3">
                        <label
                            for="exampleform"
                            class="form-label">Hari diundi</label
                        >
                        <input
                            bind:value="{pasaran_pasarandiundi_field}"
                            type="text"
                            class="form-control required"
                            placeholder="Hari diundi"
                            aria-label="Hari diundi"
                        />
                    </div>
                </Col>
                <Col xs="6">
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label"
                            >Tutup</label
                        >
                        <input
                            bind:value="{pasaran_jamtutup_field}"
                            on:keyup={handleKeyboard_time}
                            minlength="0"
                            maxlength="8"
                            type="text"
                            style="text-align:center;"
                            class="form-control required"
                            placeholder="Tutup"
                            aria-label="Tutup"
                        />
                    </div>
                    <div class="mb-3">
                        <label
                            for="exampleFormControlInput1"
                            class="form-label">Jadwal</label
                        >
                        <input
                            bind:value="{pasaran_jamjadwal_field}"
                            on:keyup={handleKeyboard_time}
                            minlength="0"
                            maxlength="8"
                            type="text"
                            style="text-align:center;"
                            class="form-control required"
                            placeholder="Jadwal"
                            aria-label="Jadwal"
                        />
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label"
                            >Buka</label
                        >
                        <input
                            bind:value="{pasaran_jamopen_field}"
                            on:keyup={handleKeyboard_time}
                            minlength="0"
                            maxlength="8"
                            type="text"
                            style="text-align:center;"
                            class="form-control required"
                            placeholder="Buka"
                            aria-label="Buka"
                        />
                    </div>
                    <div class="mb-3">
                        <label
                            for="exampleForm"
                            class="form-label">Type</label
                        >
                        <select
                            bind:value="{pasaran_tipepasaran_field}"
                            class="form-control"
                        >
                            <option value=""></option>
                            <option value="UMUM">UMUM</option>
                            <option value="WAJIB">WAJIB</option>
                        </select>
                    </div>
                </Col>
            </Row>
        </Container>
    </slot:template>
    <slot:template slot="footer">
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
</Modal>
