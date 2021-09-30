document.querySelector("#convert")?.addEventListener("click", () => {
  const parser = new DOMParser();
  const root = parser.parseFromString(
    document.querySelector("#input")?.value,
    "text/html"
  );

  console.log(root);
});
