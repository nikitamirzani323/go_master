<script>
    import Home from "../invoice/Home.svelte";

    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let listInvoice = [];
    let record = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "INVOICE_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else {
            initInvoice();
        }
    }
    async function initInvoice() {
        const res = await fetch("/api/invoice", {
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
            record = json.record;
            if (record !== null) {
                let css_winlose = "color:red;font-weight:bold;";
                totalrecord = record.length;
                for (var i = 0; i < record.length; i++) {
                    if(parseInt(record[i]["invoice_winlose"]) > 0){
                        css_winlose = "color:blue;font-weight:bold;";
                    }else{
                        css_winlose = "color:red;font-weight:bold;";
                    }
                    listInvoice = [
                        ...listInvoice,
                        {
                            invoice_no: i+1,
                            invoice_id: record[i]["invoice_id"],
                            invoice_company: record[i]["invoice_company"],
                            invoice_date: record[i]["invoice_date"],
                            invoice_name: record[i]["invoice_name"],
                            invoice_winlose: record[i]["invoice_winlose"],
                            invoice_winlose_css: css_winlose,
                            invoice_status: record[i]["invoice_status"],
                            invoice_statuscss: record[i]["invoice_statuscss"],
                        },
                    ];
                }
            } else {
                // alert("Error");
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listInvoice = [];
        totalrecord = 0;
        setTimeout(function () {
            initInvoice();
        }, 1000);
    };
    initapp();
</script>
<Home 
    on:handleRefreshData={handleRefreshData}
    {token}
    {master}
    {totalrecord}
    {listInvoice} />