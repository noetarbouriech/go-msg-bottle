// Enable wrapping
import {wrap} from 'svelte-spa-router/wrap'

// Components
import Login from './routes/Login.svelte'
import Home from './routes/Home.svelte'
import NotFound from './routes/NotFound.svelte'

// Route definition
export default {
    '/': Login,
    '/app': wrap({
        component: Home,
        conditions: [
            (detail) => {
                let jwt_cookie = document.cookie.match(/^(.*;)?\s*jwt\s*=\s*[^;]+(.*)?$/);
                // return jwt_cookie !== null; // Doesn't work since cookie is HTTPOnly
                return true
            }
        ]
    }),

    // Catch-all
    '*': NotFound,
}
