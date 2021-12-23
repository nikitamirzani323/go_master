<script>
    import Home from "../setting/Home.svelte";

    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    let maintenance_start_field = ""
    let maintenance_end_field = ""
    async function initapp() {
        const res = await fetch("/api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "SETTING_HOME",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else {
            initSetting();
        }
    }
    async function initSetting() {
        const res = await fetch("/api/setting", {
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
        let record = json.record;
        if (json.status === 400) {
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                maintenance_start_field = record[i]["maintenance_start"];
                maintenance_end_field = record[i]["maintenance_end"];
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp()
</script>
<Home
    {token}
    {master}
    {maintenance_start_field}
    {maintenance_end_field}/>