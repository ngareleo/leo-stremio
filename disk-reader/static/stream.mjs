const env = (() => {
    const [, route, id] = window.location.pathname.split("/");
    return route == "open" ? { route, id } : null;
})();

if (env == null) {
    throw Error("problem loading executing env");
}

async function pingStreamServer() {
    const { id } = env;
    let response;
    try {
        response = await fetch(`${window.location.origin}/stream/${id}`);
    } catch (e) {
        console.error("something went wrong");
    }
    const message = await response.text();
    console.log("message from server ", message);
}

async function main() {
    await pingStreamServer();
}

main()
    .then(() => {
        console.log("execution finished");
    })
    .catch((e) => {
        console.log("problem during execution");
        console.error(e);
    });
