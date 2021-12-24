<script>
    import { Row, Col, Container } from "sveltestrap";
    let token = localStorage.getItem("token");
    let master = localStorage.getItem("master");
    async function initapp() {
        const res = await fetch("/api/init", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                master: master,
                page: "ADMIN_HOME",
            }),
        });
        const json = await res.json();
        console.log(json)
        if (json.status === 400) {
            localStorage.clear();
            window.location.href = "/";
        }
    }
    initapp();
</script>

<Container fluid style="margin-top: 70px;">
    <Row>
        <Col xs="4">asd</Col>
    </Row>
</Container>
