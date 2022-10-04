<script>
    import { Router, Route, Link } from "svelte-navigator";
    import Home from 'src/pages/Home.svelte'
</script>


<Router>
    <Route path="/">
        <Home />
    </Route>
    <!-- <Route path="about" component={About} />
    <Route path="blog/*">
        <Route path="/">
            <Blog />
        </Route>
        <Route path=":id" component={BlogPost} />
    </Route>
    <Route path="search/:query" let:params>
        <Search query={params.query} />
    </Route> -->
</Router>