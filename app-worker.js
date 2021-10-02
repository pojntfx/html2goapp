const cacheName = "app-" + "8ec9d4cf3d51497b2da2ac90fa0270264f7e07fa";

self.addEventListener("install", event => {
  console.log("installing app worker 8ec9d4cf3d51497b2da2ac90fa0270264f7e07fa");

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
          "/html2goapp/web/default.png",
          "/html2goapp/web/large.png",
          "https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly-addons.css",
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
  console.log("app worker 8ec9d4cf3d51497b2da2ac90fa0270264f7e07fa is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
