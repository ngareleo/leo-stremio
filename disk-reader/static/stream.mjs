const env = (() => {
    const [, route, id] = window.location.pathname.split("/");
    return route == "open" ? { route, id } : null;
})();

if (env == null) {
    throw Error("problem loading executing env");
}

async function pingStreamServer() {
    const { route, id } = env;
    let message;
    try {
        message = await fetch(`${window.location.origin}/stream/${id}`);
    } catch (e) {
        console.error("something went wrong");
    }
    console.log("message from server ", message);
}

async function main() {
    console.log(window.location);
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
