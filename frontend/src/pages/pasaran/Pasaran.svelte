<script>
    import Home from "../pasaran/Home.svelte";
    import Entry from "../pasaran/Entry.svelte";

    let listPasaran = [];
    let record = "";
    let totalrecord = 0;
    let sData = "";
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let idpasarantogel = "";
    let pasaran_nmpasarantogel_field = "";
    let pasaran_tipepasaran_field = "";
    let pasaran_urlpasaran_field = "";
    let pasaran_pasarandiundi_field = "";
    let pasaran_jamtutup_field = "";
    let pasaran_jamjadwal_field = ""
    let pasaran_jamopen_field = "";
    let pasaran_create_field = "";
    let pasaran_update_field = "";
    const handleRefreshData = (e) => {
        listPasaran = [];
        totalrecord = 0;
        setTimeout(function () {
            initPasaran();
        }, 1000);
    };
    const handleBackHalaman = () => {
        idpasarantogel = "";
        sData = "";
        listPasaran = [];
        handleRefreshData("all");
    };
    const handleEditData = (e) => {
        idpasarantogel = e.detail.e;
        sData = "Edit";
        editPasaran(idpasarantogel);
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
                page: "PASARAN_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else {
            // setTimeout(function(){ initPasaran() }, 1000);
            initPasaran();
        }
    }
    async function initPasaran() {
        const res = await fetch("/api/allpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let no = 0;
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1
                    listPasaran = [
                        ...listPasaran,
                        {
                            pasaran_no: no,
                            pasaran_idpasarantogel: record[i]["pasaran_idpasarantogel"],
                            pasaran_nmpasarantogel: record[i]["pasaran_nmpasarantogel"],
                            pasaran_tipepasaran: record[i]["pasaran_tipepasaran"],
                            pasaran_urlpasaran: record[i]["pasaran_urlpasaran"],
                            pasaran_pasarandiundi: record[i]["pasaran_pasarandiundi"],
                            pasaran_jamtutup: record[i]["pasaran_jamtutup"],
                            pasaran_jamjadwal: record[i]["pasaran_jamjadwal"],
                            pasaran_jamopen: record[i]["pasaran_jamopen"],
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
    async function editPasaran(e) {
        const res = await fetch("/api/pasarandetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                pasarancode: e,
                master: master,
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status === 400) {
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                pasaran_nmpasarantogel_field = record[i]["pasaran_nmpasarantogel"];
                pasaran_tipepasaran_field = record[i]["pasaran_tipepasaran"];
                pasaran_urlpasaran_field = record[i]["pasaran_urlpasaran"];
                pasaran_pasarandiundi_field = record[i]["pasaran_pasarandiundi"];
                pasaran_jamtutup_field = record[i]["pasaran_jamtutup"];
                pasaran_jamjadwal_field = record[i]["pasaran_jamjadwal"];
                pasaran_jamopen_field = record[i]["pasaran_jamopen"];
                pasaran_create_field = record[i]["pasaran_create"];
                pasaran_update_field = record[i]["pasaran_update"];
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp();
</script>

{#if sData == "" && idpasarantogel == ""}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleEditData={handleEditData}
        {totalrecord}
        {listPasaran}
        {master}
        {token}
    />
{:else}
    <Entry
        on:handleBackHalaman={handleBackHalaman}
        {sData}
        {master}
        {token}
        {idpasarantogel}
        {pasaran_nmpasarantogel_field}
        {pasaran_tipepasaran_field}
        {pasaran_urlpasaran_field}
        {pasaran_pasarandiundi_field}
        {pasaran_jamtutup_field}
        {pasaran_jamjadwal_field}
        {pasaran_jamopen_field}
        {pasaran_create_field}
        {pasaran_update_field}
    />
{/if}