const cacheName = "app-" + "cbeaa64bc8abec5e31912d95ca0d6301af346efc";

self.addEventListener("install", event => {
  console.log("installing app worker cbeaa64bc8abec5e31912d95ca0d6301af346efc");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/html2goapp",
          "/html2goapp/app.css",
          "/html2goapp/app.js",
          "/html2goapp/manifest.webmanifest",
          "/html2goapp/wasm_exec.js",
          "/html2goapp/web/app.wasm",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          "https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly.css",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker cbeaa64bc8abec5e31912d95ca0d6301af346efc is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
