<script>
  import { link } from "svelte-spa-router";
  import { Button } from "sveltestrap";
  async function handleLogout() {
    localStorage.clear();
    window.href = "/";
  }
</script>

<header>
  <!-- Fixed navbar -->
  <nav
    class="navbar navbar-expand-md navbar-dark fixed-top bg-dark"
    style="background-color: #1f2937!important;"
  >
    <div class="container-fluid">
      <a class="navbar-brand" href="/" style="font-size: 13px;">MASTER SDSB4D</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarCollapse"
        aria-controls="navbarCollapse"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon" />
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" style="font-size: 13px;" use:link href="/"
              >DASHBOARD</a
            >
          </li>
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" style="font-size: 13px;" use:link href="/invoice"
              >INVOICE</a
            >
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              id="navbarDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
              style="font-size: 13px;">
              COMPANY
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <a class="dropdown-item" style="font-size: 13px;" use:link href="/company">COMPANY</a>
              </li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              id="navbarDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false" 
              style="font-size: 13px;" 
            >
              MASTER
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <a class="dropdown-item" style="font-size: 13px;" use:link href="/pasaran">PASARAN</a>
                <a class="dropdown-item" style="font-size: 13px;" use:link href="/domain">DOMAIN</a>
              </li>
            </ul>
          </li>
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" style="font-size: 13px;" use:link href="/setting">SETTING</a>
          </li>
        </ul>
        <form class="d-flex">
          <button 
            on:click={() => {
              handleLogout();
            }}
            class="btn btn-warning btn-sm">LOGOUT</button>
        </form>
      </div>
    </div>
  </nav>
</header>
