function getFileId() {
    const [, id] = window.location.pathname.split("/");
    return id;
}

console.log("File Id ", getFileId());
