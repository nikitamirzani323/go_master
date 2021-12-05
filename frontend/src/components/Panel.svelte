<script>
    export let height_body = "650px";
    export let card_header = true
    export let card_footer = true
    export let css_body = "overflow:scroll;";
    export let css_footer = "";
</script>

<div class="card" style="border-radius: 0px;margin-top:10px;">
    {#if card_header == true}
    <div class="card-header" style="">
        <slot name="cheader" />
    </div>
    <slot name="csearch" />
    {/if}
    <div class="card-body" style="height:{height_body};{css_body}">
        <slot name="cbody" />
    </div>
    {#if card_footer == true}
    <div class="card-footer" style={css_footer}>
        <slot name="cfooter" />
    </div>
    {/if}
</div>
