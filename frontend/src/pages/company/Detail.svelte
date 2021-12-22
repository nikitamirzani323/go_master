<script>
    import { createEventDispatcher } from "svelte";
    import Loader from "../../components/Loader.svelte";
    import Panel from "../../components/Panel.svelte";
    import Modal from "../../components/Modal.svelte";

    export let sData = ""
    export let token = ""
    export let master = ""
    export let idcompany = ""
    export let listPasaran = []
    let listPeriode = []
    let listInvoiceMember = []
    let listInvoiceMembertemp = []
    let listInvoiceGroupPermainan = []
    let listInvoicelistbet = []
    let record = ""
    let totalrecord = 0
    let totalbet = 0
    let totalbayar = 0
    let totalcancel = 0
    let totalwin = 0
    let totalwinlose = 0
    let totalwinlose_css = "color:red;font-weight:bold;";
    let subtotal_bet = 0;
    let subtotal_bayar = 0;
    let subtotal_cancel = 0;
    let subtotal_win = 0;
    let subtotal_win_css = "";
    let subtotal_winlose = 0;
    let subtotal_winlose_css = "";
    let subtotal_bet_temp = 0;
    let subtotal_bayar_temp = 0;
    let subtotal_cancel_temp = 0;
    let subtotal_win_temp = 0;
    let subtotal_win_temp_css = "";
    let subtotal_winlose_temp = 0;
    let subtotal_winlose_temp_css = "";
    let css_totalwin = "";
    let css_totalwin_temp = "";
    let select_periode = ""
    let select_pasaran = ""
    let invoice = ""
    let pasaran = ""
    let member = ""
    let css_loader = "display: none;";
    let msgloader = "";
    let dispatch = createEventDispatcher();
    let searchinvoice = ""; 
    let filterinvoice = [];
    let searchlistbet = "";
    let filterlistbet = [];
    $: {
        if (searchinvoice) {
            filterinvoice = listPeriode.filter(
                (item) =>
                    item.pasaran_status
                        .toLowerCase()
                        .includes(searchinvoice.toLowerCase()) ||
                    item.pasaran_invoice
                        .toLowerCase()
                        .includes(searchinvoice.toLowerCase())
            );
        } else {
            filterinvoice = [...listPeriode];
        }
        if(searchlistbet) {
            filterlistbet = listInvoicelistbet.filter(
                (item) =>
                    item.bet_status
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase()) ||
                    item.bet_nomortogel
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase()) ||
                    item.bet_username
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase()) || 
                    item.bet_id
                        .toLowerCase()
                        .includes(searchlistbet.toLowerCase())
            );
        } else {
            filterlistbet = [...listInvoicelistbet];
        }
    }
    const BackHalaman = () => {
        dispatch("handleBackHalaman", "call");
    };
    async function generate() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        listPeriode = [];
        if (select_periode == "") {
            flag = true;
            msg += "The Periode is required\n";
        }
        if (select_pasaran == "") {
            flag = true;
            msg += "The Pasaran is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/companylistkeluaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "COMPANY_HOME",
                    company: idcompany,
                    periode: select_periode,
                    pasaran: parseInt(select_pasaran),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                record = json.record;
                totalwinlose = json.totalwinlose;
                if(totalwinlose > 0){
                    totalwinlose_css = "color:blue;font-weight:bold;";
                }else{
                    totalwinlose_css = "color:red;font-weight:bold;";
                }
                let css_totalmember = "";
                let css_totalbet = "";
                let css_totalbayar = "";
                let css_totalcancel = "";
                let css_winlose = "";
                let css_winlosetemp = "";
                let css_selisih = "";
                let css_revisi = "";
                if (record != null) {
                    totalrecord = record.length;
                    for (var i = 0; i < record.length; i++) {
                        let selisih = parseInt(record[i]["company_pasaran_winlose"]) - parseInt(record[i]["company_pasaran_winlosetemp"])
                        if (parseInt(selisih) > 0) {
                            css_selisih = "color:blue;font-weight:bold";
                        } else {
                            css_selisih = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_totalmember"]) > 0) {
                            css_totalmember = "color:blue;font-weight:bold";
                        } else {
                            css_totalmember = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_totalbet"]) > 0) {
                            css_totalbet = "color:blue;font-weight:bold";
                        } else {
                            css_totalbet = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_totaloutstanding"]) > 0) {
                            css_totalbayar = "color:blue;font-weight:bold";
                        } else {
                            css_totalbayar = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_totalcancelbet"]) > 0) {
                            css_totalcancel = "color:blue;font-weight:bold";
                        } else {
                            css_totalcancel = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_winlose"]) > 0) {
                            css_winlose = "color:blue;font-weight:bold";
                        } else {
                            css_winlose = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_winlosetemp"]) > 0) {
                            css_winlosetemp = "color:blue;font-weight:bold";
                        } else {
                            css_winlosetemp = "color:red;font-weight:bold";
                        }
                        if (parseInt(record[i]["company_pasaran_revisi"]) > 0) {
                            css_revisi = "color:blue;font-weight:bold";
                        } else {
                            css_revisi = "color:red;font-weight:bold";
                        }
                        listPeriode = [
                            ...listPeriode,
                            {
                                no: record[i]["company_pasaran_no"],
                                company_pasaran_invoice: record[i]["company_pasaran_invoice"].toString(),
                                company_pasaran_idcompp: record[i]["company_pasaran_idcompp"],
                                company_pasaran_code: record[i]["company_pasaran_code"],
                                company_pasaran_periode: record[i]["company_pasaran_periode"],
                                company_pasaran_name: record[i]["company_pasaran_name"],
                                company_pasaran_keluaran: record[i]["company_pasaran_keluaran"],
                                company_pasaran_tanggal: record[i]["company_pasaran_tanggal"],
                                company_pasaran_totalmember:record[i]["company_pasaran_totalmember"],
                                company_pasaran_totalmember_css: css_totalmember,
                                company_pasaran_totalbet: record[i]["company_pasaran_totalbet"],
                                company_pasaran_totalbet_css: css_totalbet,
                                company_pasaran_totaloutstanding:record[i]["company_pasaran_totaloutstanding"],
                                company_pasaran_totaloutstanding_css: css_totalbayar,
                                company_pasaran_totalcancelbet:record[i]["company_pasaran_totalcancelbet"],
                                company_pasaran_totalcancelbet_css:css_totalcancel,
                                company_pasaran_msgrevisi: record[i]["company_pasaran_msgrevisi"],
                                company_pasaran_revisi: record[i]["company_pasaran_revisi"],
                                company_pasaran_revisi_css: css_revisi,
                                company_pasaran_winlose: record[i]["company_pasaran_winlose"],
                                company_pasaran_winlose_css: css_winlose,
                                company_pasaran_winlosetemp: record[i]["company_pasaran_winlosetemp"],
                                company_pasaran_winlosetemp_css: css_winlosetemp,
                                company_pasaran_selisih: selisih,
                                company_pasaran_selisih_css: css_selisih,
                                company_pasaran_status: record[i]["company_pasaran_status"],
                                company_pasaran_status_css: record[i]["company_pasaran_status_css"],
                            },
                        ];
                    }
                }
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function invoice_member(e) {
        listInvoiceMember = []
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicemember", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (parseInt(record[i]["totalwin"]) > 0) {
                        css_totalwin = "color:blue;font-weight:bold";
                    } else {
                        css_totalwin = "color:red;font-weight:bold";
                    }
                    subtotal_bet = subtotal_bet + parseInt(record[i]["totalbet"])
                    subtotal_bayar = subtotal_bayar + parseInt(record[i]["totalbayar"])
                    subtotal_cancel = subtotal_cancel + parseInt(record[i]["totalcancelbet"])
                    subtotal_win = subtotal_win + parseInt(record[i]["totalwin"])
                    if(subtotal_win > 0){
                        subtotal_win_css = "color:blue;font-weight:bold";
                    }else{
                        subtotal_win_css = "color:red;font-weight:bold";
                    }
                    listInvoiceMember = [
                        ...listInvoiceMember,
                        {
                            member: record[i]["member"],
                            totalbet: record[i]["totalbet"],
                            totalbayar: record[i]["totalbayar"],
                            totalcancelbet: record[i]["totalcancelbet"],
                            totalwin: record[i]["totalwin"],
                            totalwin_css: css_totalwin,
                        },
                    ];
                }
                subtotal_winlose = parseInt(subtotal_bayar) - parseInt(subtotal_cancel) - parseInt(subtotal_win)
                if(subtotal_winlose > 0){
                    subtotal_winlose_css = "color:blue;font-weight:bold";
                }else{
                    subtotal_winlose_css = "color:red;font-weight:bold";
                }
            }
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function invoice_membertemp(e) {
        listInvoiceMembertemp = []
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicemembertemp", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (parseInt(record[i]["totalwin"]) > 0) {
                        css_totalwin_temp = "color:blue;font-weight:bold";
                    } else {
                        css_totalwin_temp = "color:red;font-weight:bold";
                    }
                    subtotal_bet_temp = subtotal_bet_temp + parseInt(record[i]["totalbet"])
                    subtotal_bayar_temp = subtotal_bayar_temp + parseInt(record[i]["totalbayar"])
                    subtotal_cancel_temp = subtotal_cancel_temp + parseInt(record[i]["totalcancelbet"])
                    subtotal_win_temp = subtotal_win_temp + parseInt(record[i]["totalwin"])
                    if(subtotal_win_temp > 0){
                        subtotal_win_temp_css = "color:blue;font-weight:bold";
                    }else{
                        subtotal_win_temp_css = "color:red;font-weight:bold";
                    }
                    listInvoiceMembertemp = [
                        ...listInvoiceMembertemp,
                        {
                            member: record[i]["member"],
                            totalbet: record[i]["totalbet"],
                            totalbayar: record[i]["totalbayar"],
                            totalcancelbet: record[i]["totalcancelbet"],
                            totalwin: record[i]["totalwin"],
                            totalwin_css: css_totalwin,
                        },
                    ];
                }
                subtotal_winlose_temp = parseInt(subtotal_bayar_temp) - parseInt(subtotal_cancel_temp) - parseInt(subtotal_win_temp)
                if(subtotal_winlose_temp > 0){
                    subtotal_winlose_temp_css = "color:blue;font-weight:bold";
                }else{
                    subtotal_winlose_temp_css = "color:red;font-weight:bold";
                }
            }
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function invoice_membersync() {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicemembersync", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice),
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
        invoice_membertemp(invoice)
    }
    async function invoice_grouppermainan(e) {
        listInvoiceGroupPermainan = []
        listInvoicelistbet = []
        member = e
        totalbet = 0
        totalbayar = 0
        totalcancel = 0
        totalwin = 0
        totalwinlose = 0
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicegrouppermainan", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice),
                username: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listInvoiceGroupPermainan = [
                        ...listInvoiceGroupPermainan,
                        {
                            permainan: record[i]["permainan"],
                        },
                    ];
                }
            }
        } else {
            msgloader = json.message;
        }
        if(e != "" && e != undefined){
            let modalbetmember = new bootstrap.Modal(
                document.getElementById("modalbetmember")
            );
            modalbetmember.show()
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function invoice_listbet(e) {
        listInvoicelistbet = []
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicelistpermainan", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice),
                permainan: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listInvoicelistbet = [
                        ...listInvoicelistbet,
                        {
                            bet_id: record[i]["bet_id"].toString(),
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: record[i]["bet_statuscss"],
                        },
                    ];
                }
            }
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function invoice_listbetusername(e) {
        listInvoicelistbet = []
        
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicelistpermainanmember", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice),
                username: member,
                permainan: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalbet = json.totalbet
            totalbayar = json.subtotal
            totalcancel = json.subtotalcancel
            totalwin = json.subtotalwin
            totalwinlose = parseInt(totalbayar) - parseInt(totalcancel) - parseInt(totalwin)
            if(totalwinlose > 0){
                totalwinlose_css = "color:blue;font-weight:bold;";
            }else{
                totalwinlose_css = "color:red;font-weight:bold;";
            }
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listInvoicelistbet = [
                        ...listInvoicelistbet,
                        {
                            bet_id: record[i]["bet_id"].toString(),
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: record[i]["bet_statuscss"],
                        },
                    ];
                }
            }
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function invoice_permainanstatus(e) {
        listInvoicelistbet = []
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/companyinvoicelistpermainanstatus", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: idcompany,
                invoice: parseInt(invoice),
                status: e,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listInvoicelistbet = [
                        ...listInvoicelistbet,
                        {
                            bet_id: record[i]["bet_id"].toString(),
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: record[i]["bet_statuscss"],
                        },
                    ];
                }
            }
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    const editKeluaran = (e,y) => {
        invoice = e
        pasaran = y
        let myModalPasaran = new bootstrap.Modal(
            document.getElementById("modalEditKeluaran")
        );
        myModalPasaran.show();
        listInvoiceMember = []
        listInvoiceGroupPermainan = []
        listInvoicelistbet = []
        totalbet = 0
        totalbayar = 0
        totalcancel = 0
        totalwin = 0
        totalwinlose = 0
        subtotal_bet = 0;
        subtotal_bayar = 0;
        subtotal_cancel = 0;
        subtotal_win = 0;
        subtotal_winlose = 0;
        subtotal_bet_temp = 0;
        subtotal_bayar_temp = 0;
        subtotal_cancel_temp = 0;
        subtotal_winlose_temp = 0;
        invoice_member(e)
        invoice_membertemp(e)
        invoice_grouppermainan("")
    };
    const handleSelectPermainan = (event) => {
        invoice_listbet(event.target.value)
    };
    const handleSelectPermainan_dua = (event) => {
        totalbet = 0
        totalbayar = 0
        totalcancel = 0
        totalwin = 0
        totalwinlose = 0
        invoice_listbetusername(event.target.value)
    };
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <button
                on:click={() => {
                    BackHalaman();
                }}
                class="btn btn-dark btn-sm"
                style="border-radius: 0px;"
            >
                Back
            </button>
        </div>
        <div class="clearfix" />
        
        <div class="col-sm-12">
            <Panel 
                card_footer={true}
                height_body="600px">
                <slot:template slot="cheader">
                    List Keluaran
                </slot:template>
                <slot:template slot="csearch">
                    <div class="row" style="padding: 5px;">
                        <div class="col-lg-3">
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
                        <div class="col-lg-3">
                            <label for="exampleForm" class="form-label">Pasaran</label>
                            <select
                                bind:value="{select_pasaran}" 
                                class="form-control">
                                {#each listPasaran as rec }
                                    <option value="{rec.company_idcomppasaran}">{rec.company_nmpasarantogel}</option>
                                {/each}
                            </select>
                        </div>
                        <div class="col-lg-2">
                            <br>
                            <button
                                on:click={() => {
                                    generate();
                                }}
                                class="btn btn-warning btn-sm"
                                style="border-radius: 0px;margin-top:10px;"
                            >
                                Generate
                            </button>
                        </div>
                    </div>
                    
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value="{searchinvoice}"
                            type="text"
                            class="form-control"
                            placeholder="Search"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="cbody">
                    {#if totalrecord > 0}
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">&nbsp;</th>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">NO</th>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">&nbsp;</th>
                                    <th
                                        width="7%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">DATE</th>
                                    <th
                                        NOWRAP
                                        width="1%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">INVOICE</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">PERIODE</th>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">PASARAN</th>
                                    <th
                                        width="5%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">KELUARAN</th>
                                    <th
                                        NOWRAP
                                        width="1%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">REVISI</th>
                                    <th
                                        width="6%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">MEMBER</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">BET</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">BAYAR</th>
                                    <th
                                        width="5%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">CANCEL</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">WINLOSE</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">WINLOSE TEMP</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;">SELISIH</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterinvoice as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                editKeluaran(
                                                    rec.company_pasaran_invoice,
                                                    rec.company_pasaran_name
                                                );
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;">
                                            <i class="bi bi-pencil"></i>
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;">{rec.no}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.company_pasaran_status_css}">{rec.company_pasaran_status}</td>

                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;">{rec.company_pasaran_tanggal}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.company_pasaran_invoice}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.company_pasaran_periode}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.company_pasaran_name}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;font-weight:bold;color:black;">{rec.company_pasaran_keluaran}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_revisi_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_revisi
                                            )}
                                            {#if rec.company_pasaran_revisi > 0} 
                                            <i class="bi bi-chat-right-text" data-bs-toggle="tooltip" data-bs-placement="top" data-bs-html="true" title="{rec.company_pasaran_msgrevisi}"></i> 
                                            {/if}
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_totalmember_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_totalmember
                                            )}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_totalbet_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_totalbet
                                            )}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_totaloutstanding_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_totaloutstanding
                                            )}</td >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_totalcancelbet_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_totalcancelbet
                                            )}</td >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_winlose_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_winlose
                                            )}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_winlosetemp_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_winlosetemp
                                            )}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.company_pasaran_selisih_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.company_pasaran_selisih
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
                    <div class="float-end">
                        TOTAL WINLOSE : <span style="{totalwinlose_css}">{new Intl.NumberFormat().format(
                                            totalwinlose
                                        )}</span>
                    </div>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>
<Modal
    modal_id={"modalEditKeluaran"}
    modal_size={"modal-fullscreen"}
    modal_body_height={"height:200px;"}
    modal_footer_flag={false}>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">
            INVOICE : {invoice}<br />
            PASARAN : {pasaran}
        </h5>
    </slot:template>
    <slot:template slot="body">
        <ul class="nav nav-pills mb-3" id="pills-tab" role="tablist">
            <li class="nav-item" role="presentation">
                <button class="nav-link active" id="pills-information-tab" data-bs-toggle="pill" data-bs-target="#pills-information" type="button" role="tab" aria-controls="pills-information" aria-selected="true">Information</button>
            </li>
            <li class="nav-item" role="presentation">
                <button class="nav-link" id="pills-member-tab" data-bs-toggle="pill" data-bs-target="#pills-member" type="button" role="tab" aria-controls="pills-member" aria-selected="true">MEMBER</button>
            </li>
            <li 
                on:click={() => {
                    invoice_grouppermainan();
                }}
                class="nav-item" role="presentation">
                <button class="nav-link" id="pills-groupbet-tab" data-bs-toggle="pill" data-bs-target="#pills-groupbet" type="button" role="tab" aria-controls="pills-groupbet" aria-selected="false">GROUP BET</button>
            </li>
            <li 
                on:click={() => {
                    invoice_grouppermainan();
                }}
                class="nav-item" role="presentation">
                <button class="nav-link" id="pills-allbet-tab" data-bs-toggle="pill" data-bs-target="#pills-allbet" type="button" role="tab" aria-controls="pills-allbet" aria-selected="false">ALL BET</button>
            </li>
            <li 
                on:click={() => {
                    invoice_permainanstatus("WINNER");
                }}
                class="nav-item" role="presentation">
                <button class="nav-link" id="pills-winner-tab" data-bs-toggle="pill" data-bs-target="#pills-winner" type="button" role="tab" aria-controls="pills-winner" aria-selected="false">WINNER</button>
            </li>
            <li 
                on:click={() => {
                    invoice_permainanstatus("CANCEL");
                }}
                class="nav-item" role="presentation">
                <button class="nav-link" id="pills-cancel-tab" data-bs-toggle="pill" data-bs-target="#pills-cancel" type="button" role="tab" aria-controls="pills-cancel" aria-selected="false">CANCEL</button>
            </li>
        </ul>
        <div class="tab-content" id="pills-tabContent">
            <div class="tab-pane fade show active" id="pills-information" role="tabpanel" aria-labelledby="pills-information-tab">
                <div class="row">
                    <div class="col-md-2" style="padding:0px;margin:0px;">
                        <div class="card" style="border-radius: 0px;margin-top:10px;">
                            <div class="card-header" style="">
                                INFORMATION 
                                <div class="float-end">
                                   
                                    <button
                                        on:click={() => {
                                            invoice_revisi();
                                        }}
                                        class="btn btn-danger btn-sm">Revisi</button>
                                </div>
                            </div>
                            <div class="card-body" style="height:350px;">
                                
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="tab-pane fade show" id="pills-member" role="tabpanel" aria-labelledby="pills-member-tab">
                <div class="row">
                    <div class="col-md-5" style="padding:0px;margin:0px;">
                        <Panel height_body="450px" card_footer={true} >
                            <slot:template slot="cheader">
                                <h2 style="font-size:14px;">TBL_KELUARAN_DETAIL</h2>
                            </slot:template>
                            <slot:template slot="cbody">
                                <table class="table" width="50%">
                                    <thead>
                                        <tr>
                                            <th width="*"
                                                style="text-align: left;vertical-align: top;font-size: 13px;">MEMBER</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL BET</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL BAYAR</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL CANCEL</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL WIN</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {#each listInvoiceMember as rec }
                                            <tr>
                                                <td 
                                                    on:click={() => {
                                                        invoice_grouppermainan(rec.member);
                                                    }}
                                                    NOWRAP
                                                    style="text-decoration:underline;cursor:pointer;text-align: left;vertical-align: top;font-size: 12px;">{rec.member}</td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalbet
                                                    )}
                                                </td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalbayar
                                                    )}
                                                </td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalcancelbet
                                                    )}
                                                </td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;{css_totalwin}">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalwin
                                                    )}
                                                </td>
                                            </tr>
                                        {/each}
                                    </tbody>
                                </table>
                            </slot:template>
                            <slot:template slot="cfooter">
                                <table style="width:100%;">
                                    <tbody>
                                        <tr>
                                            <td NOWRAP width="90%" style="text-align:right;vertical-align:top;">SUBTOTAL BET</td>
                                            <td NOWRAP width="1%" style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP width="*" style="text-align:right;vertical-align:top;font-size: 12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_bet
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL BAYAR</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_bayar
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL CANCEL</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;color:red;">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_cancel
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL WIN</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;{subtotal_win_css}">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_win
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL WINLOSE</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;{subtotal_winlose_css}">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_winlose
                                                )}
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </slot:template>
                        </Panel>
                    </div>
                    <div class="col-md-2" style="padding:0px;margin:0px;">
                        <center>
                            <button
                                on:click={() => {
                                    invoice_membersync();
                                }}
                                class="btn btn-danger">SYNC</button>
                        </center>
                    </div>
                    <div class="col-md-5" style="padding:0px;margin:0px;">
                        <Panel height_body="450px" card_footer={true} css_body="">
                            <slot:template slot="cheader">
                                <h2 style="font-size:14px;">TBL_KELUARAN_MEMBER</h2>
                            </slot:template>
                            <slot:template slot="cbody">
                                <table class="table" width="50%">
                                    <thead>
                                        <tr>
                                            <th width="*"
                                                style="text-align: left;vertical-align: top;font-size: 13px;">MEMBER</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL BET</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL BAYAR</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL CANCEL</th>
                                            <th NOWRAP
                                                width="20%"
                                                style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL WIN</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {#each listInvoiceMembertemp as rec }
                                            <tr>
                                                <td NOWRAP
                                                    style="text-align: left;vertical-align: top;font-size: 12px;">{rec.member}</td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalbet
                                                    )}
                                                </td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalbayar
                                                    )}
                                                </td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalcancelbet
                                                    )}
                                                </td>
                                                <td NOWRAP
                                                    style="text-align: right;vertical-align: top;font-size: 12px;{css_totalwin_temp}">
                                                    {new Intl.NumberFormat().format(
                                                        rec.totalwin
                                                    )}
                                                </td>
                                            </tr>
                                        {/each}
                                    </tbody>
                                </table>
                            </slot:template>
                            <slot:template slot="cfooter">
                                <table style="width:100%;">
                                    <tbody>
                                        <tr>
                                            <td NOWRAP width="90%" style="text-align:right;vertical-align:top;">SUBTOTAL BET</td>
                                            <td NOWRAP width="1%" style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP width="*" style="text-align:right;vertical-align:top;font-size: 12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_bet_temp
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL BAYAR</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_bayar_temp
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL CANCEL</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;color:red;">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_cancel_temp
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL WIN</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;{subtotal_win_temp_css}">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_win_temp
                                                )}
                                            </td>
                                        </tr>
                                        <tr>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">SUBTOTAL WINLOSE</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;">:</td>
                                            <td NOWRAP style="text-align:right;vertical-align:top;font-size: 12px;{subtotal_winlose_temp_css}">
                                                {new Intl.NumberFormat().format(
                                                    subtotal_winlose_temp
                                                )}
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </slot:template>
                        </Panel>
                    </div>
                </div>
            </div>
            <div class="tab-pane fade" id="pills-groupbet" role="tabpanel" aria-labelledby="pills-groupbet-tab">
                <Panel height_body="450px" card_footer="false" css_body="">
                    <slot:template slot="cheader">
                        GROUP BET
                    </slot:template>
                    <slot:template slot="cbody">
                        <table>
                            <thead>
                                <tr>
                                    <th width="*"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each listInvoiceGroupPermainan as rec }
                                    <tr>
                                        <td NOWRAP
                                            style="text-decoration:underline;cursor:pointer;text-align: left;vertical-align: top;font-size: 12px;">{rec.permainan}</td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </Panel>
            </div>
            <div class="tab-pane fade" id="pills-allbet" role="tabpanel" aria-labelledby="pills-allbet-tab">
                <Panel height_body="450px" card_header="false" card_footer="false" css_body="">
                    <slot:template slot="cbody">
                        <div class="row">
                            <div class="col-lg-5" >
                                <select
                                    on:change={handleSelectPermainan}
                                    class="form-control">
                                    <option value="">--Select--</option>
                                    {#each listInvoiceGroupPermainan as rec }
                                        <option value="{rec.permainan}">{rec.permainan}</option>
                                    {/each}
                                </select>
                            </div>
                            <div class="col-lg-7" >
                                <input
                                    bind:value="{searchlistbet}"
                                    type="text"
                                    class="form-control"
                                    placeholder="Search"
                                    aria-label="Search"/>
                            </div>
                        </div>
                        <div class="clearfix"></div><br>
                        <table class="table" width="100%">
                            <thead>
                                <tr>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                    <th
                                        width="1%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterlistbet as rec}
                                <tr>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}"
                                        >{rec.bet_status}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_id}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_datetime}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_username}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_ipaddress}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_device}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_timezone}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_typegame}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_nomortogel}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_bet
                                        )}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                        >{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                        >{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_bayar
                                        )}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_win}x</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_totalwin
                                        )}</td>
                                </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </Panel>
            </div>
            <div class="tab-pane fade" id="pills-winner" role="tabpanel" aria-labelledby="pills-winner-tab">
                <Panel height_body="450px" card_header="false" card_footer="false" css_body="">
                    <slot:template slot="cbody">
                        <div class="row">
                            
                            <div class="col-lg-7" >
                                <input
                                    bind:value="{searchlistbet}"
                                    type="text"
                                    class="form-control"
                                    placeholder="Search"
                                    aria-label="Search"/>
                            </div>
                        </div>
                        <div class="clearfix"></div><br>
                        <table class="table" width="100%">
                            <thead>
                                <tr>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                    <th
                                        width="1%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterlistbet as rec}
                                <tr>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}"
                                        >{rec.bet_status}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_id}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_datetime}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_username}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_ipaddress}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_device}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_timezone}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_typegame}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_nomortogel}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_bet
                                        )}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                        >{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                        >{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_bayar
                                        )}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_win}x</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_totalwin
                                        )}</td>
                                </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </Panel>
            </div>
            <div class="tab-pane fade" id="pills-cancel" role="tabpanel" aria-labelledby="pills-cancel-tab">
                <Panel height_body="450px" card_header="false" card_footer="false" css_body="">
                    <slot:template slot="cbody">
                        <div class="row">
                            
                            <div class="col-lg-7" >
                                <input
                                    bind:value="{searchlistbet}"
                                    type="text"
                                    class="form-control"
                                    placeholder="Search"
                                    aria-label="Search"/>
                            </div>
                        </div>
                        <div class="clearfix"></div><br>
                        <table class="table" width="100%">
                            <thead>
                                <tr>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                    <th
                                        width="1%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterlistbet as rec}
                                <tr>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}"
                                        >{rec.bet_status}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_id}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_datetime}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_username}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_ipaddress}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_device}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_timezone}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_typegame}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: left;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_nomortogel}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_bet
                                        )}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                        >{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                        >{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_bayar
                                        )}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;"
                                        >{rec.bet_win}x</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                        >{new Intl.NumberFormat().format(
                                            rec.bet_totalwin
                                        )}</td>
                                </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                </Panel>
            </div>
        </div>
    </slot:template>
</Modal>
<Modal
    modal_id={"modalbetmember"}
    modal_size={"modal-dialog-centered modal-xl"}
    modal_body_height={"height:500px;overflow:scroll;"}
    modal_headermiddle_flag={true}>
    modal_footer_flag={true}>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">
            INVOICE : {invoice}<br />
            MEMBER : {member}
        </h5>
    </slot:template>
    <slot:template slot="headermiddle">
        <div class="row">
            <div class="col-lg-5" >
                <select
                    on:change={handleSelectPermainan_dua}
                    class="form-control">
                    <option value="">--Select--</option>
                    {#each listInvoiceGroupPermainan as rec }
                        <option value="{rec.permainan}">{rec.permainan}</option>
                    {/each}
                </select>
            </div>
            <div class="col-lg-7" >
                <input
                    bind:value="{searchlistbet}"
                    type="text"
                    class="form-control"
                    placeholder="Search"
                    aria-label="Search"/>
            </div>
        </div>
    </slot:template>
    <slot:template slot="body">
        <table class="table" width="100%">
            <thead>
                <tr>
                    <th
                        width="1%"
                        style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                    <th
                        width="1%"
                        style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                    <th
                        width="10%"
                        style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                    <th
                        width="*"
                        style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                    <th
                        width="1%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                    <th
                        width="20%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                    <th
                        width="10%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                    <th
                        width="10%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                    <th
                        width="20%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                    <th
                        width="7%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                    <th
                        width="7%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                </tr>
            </thead>
            <tbody>
                {#each filterlistbet as rec}
                <tr>
                    <td
                        NOWRAP
                        style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}"
                        >{rec.bet_status}</td>
                    <td
                        NOWRAP
                        style="text-align: center;vertical-align: top;font-size: 12px;"
                        >{rec.bet_id}</td>
                    <td
                        NOWRAP
                        style="text-align: center;vertical-align: top;font-size: 12px;"
                        >{rec.bet_datetime}</td>
                    <td
                        NOWRAP
                        style="text-align: left;vertical-align: top;font-size: 12px;"
                        >{rec.bet_username}</td>
                    <td
                        NOWRAP
                        style="text-align: left;vertical-align: top;font-size: 12px;"
                        >{rec.bet_ipaddress}</td>
                    <td
                        NOWRAP
                        style="text-align: left;vertical-align: top;font-size: 12px;"
                        >{rec.bet_device}</td>
                    <td
                        NOWRAP
                        style="text-align: left;vertical-align: top;font-size: 12px;"
                        >{rec.bet_timezone}</td>
                    <td
                        NOWRAP
                        style="text-align: left;vertical-align: top;font-size: 12px;"
                        >{rec.bet_typegame}</td>
                    <td
                        NOWRAP
                        style="text-align: left;vertical-align: top;font-size: 12px;"
                        >{rec.bet_nomortogel}</td>
                    <td
                        NOWRAP
                        style="text-align: right;vertical-align: top;font-size: 12px;"
                        >{new Intl.NumberFormat().format(
                            rec.bet_bet
                        )}</td>
                    <td
                        NOWRAP
                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                        >{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                    <td
                        NOWRAP
                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                        >{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                    <td
                        NOWRAP
                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                        >{new Intl.NumberFormat().format(
                            rec.bet_bayar
                        )}</td>
                    <td
                        NOWRAP
                        style="text-align: right;vertical-align: top;font-size: 12px;"
                        >{rec.bet_win}x</td>
                    <td
                        NOWRAP
                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                        >{new Intl.NumberFormat().format(
                            rec.bet_totalwin
                        )}</td>
                </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
    <slot:template slot="footer">
        <table>
            <tbody>
                <tr>
                    <td NOWRAP width="10%" style="text-align: right;vertical-align: top;font-size:13px;">TOTAL BET</td>
                    <td NOWRAP width="1%" style="text-align: right;vertical-align: top;font-size:13px;">:</td>
                    <td NOWRAP width="*" style="text-align: right;vertical-align: top;font-size:13px;color:blue;font-weight:bold;">
                        {new Intl.NumberFormat().format(
                            totalbet
                        )}
                    </td>
                </tr>
                <tr>
                    <td NOWRAP width="10%" style="text-align: right;vertical-align: top;font-size:13px;">TOTAL BAYAR</td>
                    <td NOWRAP width="1%" style="text-align: right;vertical-align: top;font-size:13px;">:</td>
                    <td NOWRAP width="*" style="text-align: right;vertical-align: top;font-size:13px;color:blue;font-weight:bold;">
                        {new Intl.NumberFormat().format(
                            totalbayar
                        )}
                    </td>
                </tr>
                <tr>
                    <td NOWRAP width="10%" style="text-align: right;vertical-align: top;font-size:13px;">TOTAL CANCEL</td>
                    <td NOWRAP width="1%" style="text-align: right;vertical-align: top;font-size:13px;">:</td>
                    <td NOWRAP width="*" style="text-align: right;vertical-align: top;font-size:13px;color:blue;font-weight:bold;">
                        {new Intl.NumberFormat().format(
                            totalcancel
                        )}
                    </td>
                </tr>
                <tr>
                    <td NOWRAP width="10%" style="text-align: right;vertical-align: top;font-size:13px;">TOTAL WIN</td>
                    <td NOWRAP width="1%" style="text-align: right;vertical-align: top;font-size:13px;">:</td>
                    <td NOWRAP width="*" style="text-align: right;vertical-align: top;font-size:13px;color:red;font-weight:bold;">
                        {new Intl.NumberFormat().format(
                            totalwin
                        )}
                    </td>
                </tr>
                <tr>
                    <td NOWRAP width="10%" style="text-align: right;vertical-align: top;font-size:13px;">TOTAL WINLOSE</td>
                    <td NOWRAP width="1%" style="text-align: right;vertical-align: top;font-size:13px;">:</td>
                    <td NOWRAP width="*" style="text-align: right;vertical-align: top;font-size:13px;{totalwinlose_css}">
                        {new Intl.NumberFormat().format(
                            totalwinlose
                        )}
                    </td>
                </tr>
            </tbody>
        </table>
    </slot:template>
</Modal>