<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navigation from "./components/Navigation.svelte";
	import Footer from "./components/Footer.svelte";
	import Login from "./pages/Login.svelte";
	import Home from "./pages/home/Home.svelte";
	import Company from "./pages/company/Company.svelte";
	import Pasaran from "./pages/pasaran/Pasaran.svelte";
	import Invoice from "./pages/invoice/Invoice.svelte";
	import Setting from "./pages/setting/Setting.svelte";
	import Domain from "./pages/domain/Domain.svelte";
	import NotFound from "./pages/NotFound.svelte";
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				component: Home,
			}),
			"/domain": wrap({
				component: Domain,
			}),
			"/invoice": wrap({
				component: Invoice,
			}),
			"/company": wrap({
				component: Company,
			}),
			"/pasaran": wrap({
				component: Pasaran,
			}),
			"/setting": wrap({
				component: Setting,
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navigation />
{/if}
<Router {routes} />
<Footer />
