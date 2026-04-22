<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Makler</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootswatch@5.3.3/dist/simplex/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
    <link rel="stylesheet" href="/static/css/main.css">
    <style>
      .scrollable-menu { max-height: 400px; overflow-y: auto; }
    </style>
    <script>
      /* Replace broken images with a party-coloured SVG badge */
      (function () {
        var PARTIES = [
          { key: 'CCM',     bg: '#2e7d32' },
          { key: 'CHADEMA', bg: '#29b6f6' },
          { key: 'ACT',     bg: '#7b1fa2' },
          { key: 'CUF',     bg: '#ffc107' },
          { key: 'NCCR',    bg: '#1a237e' }
        ];

        function partyInfo(name) {
          var u = (name || '').toUpperCase();
          for (var i = 0; i < PARTIES.length; i++) {
            if (u.indexOf(PARTIES[i].key) !== -1) return PARTIES[i];
          }
          return { bg: '#6c757d' };
        }

        function makeSVG(bg) {
          return 'data:image/svg+xml,' + encodeURIComponent(
            '<svg xmlns="http://www.w3.org/2000/svg" width="200" height="200" viewBox="0 0 200 200">' +
            '<rect width="200" height="200" fill="' + bg + '"/>' +
            '<circle cx="100" cy="78" r="42" fill="rgba(255,255,255,0.35)"/>' +
            '<ellipse cx="100" cy="195" rx="68" ry="50" fill="rgba(255,255,255,0.35)"/>' +
            '</svg>'
          );
        }

        document.addEventListener('error', function (e) {
          var img = e.target;
          if (img.tagName !== 'IMG' || img.dataset.fallback) return;
          img.dataset.fallback = '1';
          var p = partyInfo(img.dataset.party);
          img.src = makeSVG(p.bg);
        }, true);
      }());
    </script>
  </head>
  <body>
    {{.NavBar}}
    {{.MainContent}}
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
  </body>
</html>
