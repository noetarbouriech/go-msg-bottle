<main>
  <h1>Test</h1>
  <Router {routes} on:conditionsFailed={goToLogin}/>
</main>

<script>
import Router from 'svelte-spa-router'
import {replace} from 'svelte-spa-router'
import routes from './routes'

function goToLogin() {
  replace('/');
}
</script>

<style>
</style>
