<script>
    import Home from "../company/Home.svelte";
    import Entry from "../company/Entry.svelte";
    import Detail from "../company/Detail.svelte";

    let listCompany = [];
    let listPasaran = [];
    let record = "";
    let totalrecord = 0;
    let sData = "";
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let idcompany = "";
    let company_name_field = "";
    let company_url_field = "";
    let company_status_field = "";
    let company_create_field = "";
    let company_update_field = "";
    const handleRefreshData = (e) => {
        listCompany = [];
        totalrecord = 0;
        setTimeout(function () {
            initCompany();
        }, 1000);
    };
    const handleBackHalaman = () => {
        idcompany = "";
        sData = "";
        listCompany = [];
        handleRefreshData("all");
    };
    const handleEditData = (e) => {
        idcompany = e.detail.e;
        sData = "Edit";
        editCompany(idcompany);
    };
    const handleDetailData = (e) => {
        idcompany = e.detail.e;
        sData = "Detail";
        initPasaran(idcompany);
    };
    async function initapp() {
        const res = await fetch("/api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "COMPANY_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else {
            // setTimeout(function(){ initPasaran() }, 1000);
            initCompany();
        }
    }
    async function initCompany() {
        const res = await fetch("/api/company", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_search: "",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    let css_winlose = "color:red;font-weight:bold;";
                    let css_winlose_temp = "color:red;font-weight:bold;";
                    let css_winlose_selisih = "color:red;font-weight:bold;";
                    let selisihwinlose = parseInt(record[i]["company_winlosetemp"]) - parseInt(record[i]["company_winlose"])
                    if(parseInt(record[i]["company_winlose"]) > 0){
                        css_winlose = "color:blue;font-weight:bold;";
                    }else{
                        css_winlose = "color:red;font-weight:bold;";
                    }
                    if(parseInt(record[i]["company_winlosetemp"]) > 0){
                        css_winlose_temp = "color:blue;font-weight:bold;";
                    }else{
                        css_winlose_temp = "color:red;font-weight:bold;";
                    }
                    if(selisihwinlose){
                        css_winlose_selisih = "color:blue;font-weight:bold;";
                    }
                    listCompany = [
                        ...listCompany,
                        {
                            no: record[i]["company_no"],
                            company_idcompany: record[i]["company_idcompany"],
                            company_startjoin: record[i]["company_startjoin"],
                            company_endjoin: record[i]["company_endjoin"],
                            company_curr: record[i]["company_curr"],
                            company_name: record[i]["company_name"],
                            company_owner: record[i]["company_owner"],
                            company_phone: record[i]["company_phone"],
                            company_email: record[i]["company_email"],
                            company_periode: record[i]["company_periode"],
                            company_winlose: record[i]["company_winlose"],
                            company_winlosetemp: record[i]["company_winlosetemp"],
                            company_winlosecss: css_winlose,
                            company_winlosetempcss: css_winlose_temp,
                            company_selisihwinlose: selisihwinlose,
                            company_selisihwinlosecss: css_winlose_selisih,
                            company_status: record[i]["company_status"],
                            company_statuscss: record[i]["company_statuscss"],
                        },
                    ];
                }
            } else {
                alert("Error");
            }
        } else {
            // logout();
        }
    }
    async function editCompany(e) {
        const res = await fetch("/api/companydetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: e,
                page: "COMPANY_SAVE",
                master: master,
                sData: "Edit",
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status === 400) {
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                company_name_field = record[i]["company_name"];
                company_url_field = record[i]["company_url"];
                company_status_field = record[i]["company_status"];
                company_create_field = record[i]["company_create"];
                company_update_field = record[i]["company_update"];
            }
        }
    }
    async function initPasaran(e) {
        listPasaran = []
        const res = await fetch("/api/companylistpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company: e,
                master: master,
                page: "COMPANY_HOME",
                sdata: "New",
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listPasaran = [
                        ...listPasaran,
                        {
                            company_idcomppasaran: record[i]["company_pasaran_idcomppasaran"],
                            company_nmpasarantogel: record[i]["company_pasaran_nmpasarantogel"],
                        },
                    ];
                }
            } else {
                alert("Error");
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp();
</script>

{#if sData == "" && idcompany == ""}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleEditData={handleEditData}
        on:handleDetailData={handleDetailData}
        {totalrecord}
        {listCompany}
        {master}
        {token}
    />
{:else}
    {#if sData == "Edit" }
        <Entry
            on:handleBackHalaman={handleBackHalaman}
            {sData}
            {master}
            {token}
            {idcompany}
            {company_name_field}
            {company_url_field}
            {company_status_field}
            {company_create_field}
            {company_update_field}
        />
    {:else}
        <Detail
            on:handleBackHalaman={handleBackHalaman}
            {sData}
            {master}
            {token}
            {idcompany}
            {listPasaran}
        />
    {/if}
{/if}
