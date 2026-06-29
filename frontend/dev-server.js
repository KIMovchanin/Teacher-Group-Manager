// dev-server.js — локальный сервер: отдаёт фронтенд и проксирует API на бэкенд друга.
// Это инфраструктура, НЕ материал экзамена. Понимать построчно не обязательно.

const http = require('http');
const fs = require('fs');
const path = require('path');

const PORT = 3000;                        // порт нашего фронтенда
const BACKEND = 'http://localhost:8080';  // адрес бэкенда друга
const PUBLIC = path.join(__dirname, 'public'); // папка с HTML/CSS/JS

// пути, которые считаем "запросом к API" и пересылаем на бэкенд:
const API_PREFIXES = ['/students', '/groups', '/health'];

const MIME = {
  '.html': 'text/html; charset=utf-8',
  '.css':  'text/css; charset=utf-8',
  '.js':   'text/javascript; charset=utf-8',
};

const server = http.createServer(async (req, res) => {
  const url = req.url;

  // 1) Запрос к API -> проксируем на бэкенд друга
  if (API_PREFIXES.some(p => url === p || url.startsWith(p + '/') || url.startsWith(p + '?'))) {
    try {
      const chunks = [];
      for await (const c of req) chunks.push(c);      // собираем тело запроса (для POST)
      const body = chunks.length ? Buffer.concat(chunks) : undefined;

      const backendRes = await fetch(BACKEND + url, {
        method: req.method,
        headers: { 'content-type': req.headers['content-type'] || 'application/json' },
        body: (req.method === 'GET' || req.method === 'HEAD') ? undefined : body,
      });

      const text = await backendRes.text();
      res.writeHead(backendRes.status, {
        'content-type': backendRes.headers.get('content-type') || 'application/json',
      });
      res.end(text);
    } catch (e) {
      res.writeHead(502, { 'content-type': 'application/json' });
      res.end(JSON.stringify({ error: 'backend unavailable: ' + e.message }));
    }
    return;
  }

  // 2) Иначе отдаём статический файл из папки public
  const filePath = path.join(PUBLIC, url === '/' ? 'index.html' : url.split('?')[0]);
  fs.readFile(filePath, (err, data) => {
    if (err) {
      res.writeHead(404, { 'content-type': 'text/plain; charset=utf-8' });
      res.end('404: ' + url);
      return;
    }
    const ext = path.extname(filePath);
    res.writeHead(200, { 'content-type': MIME[ext] || 'application/octet-stream' });
    res.end(data);
  });
});

server.listen(PORT, () => {
  console.log(`Frontend:  http://localhost:${PORT}`);
  console.log(`Proxy   -> ${BACKEND}`);
});
