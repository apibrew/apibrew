window.onload = function () {
    //<editor-fold desc="Changeable Configuration Block">

    // the following lines will be replaced by docker/configurator, when it runs in a docker-container

    const config = {
        urls: [
            {
                name: "User API",
                url: "/docs/openapi.json",
            },
            {
                name: "Meta API",
                url: "/docs/openapi.json?mode=meta",
            },
            {
                name: "Internal API",
                url: "/docs/openapi.json?mode=internal",
            },
            {
                name: "Full API",
                url: "/docs/openapi.json?mode=full",
            }
        ],
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
            SwaggerUIBundle.presets.apis,
            SwaggerUIStandalonePreset
        ],
        plugins: [
            SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
    }

    let params = (new URL(document.location)).searchParams;

    if (params.has('url')) {
        config.urls = undefined
        config.url = params.get('url')
    }

    if (params.has('mode') && params.get('mode') === 'inline') {
        window.document.head.innerHTML += '<link rel="stylesheet" href="./inline-mode.css">';
    }

    if (params.has('token')) {
        config.requestInterceptor = function (request) {
            request.headers['Authorization'] = 'Bearer ' + params.get('token');
            return request;
        };
    }

    window.ui = SwaggerUIBundle(config);

    //</editor-fold>
};
