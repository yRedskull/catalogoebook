# Documento para revisar antes de colocar em produção!

* ### config.go

#### remover o seguinte código:
`gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger())`

* ### webpack.config.js

#### Verificar se os seguintes atributos estão colocados:
```javascript
{
watch: false, 
mode: 'production', 
plugins: [
        new MiniCssExtractPlugin({ filename: 'css/[name].css' }),
        new WebpackObfuscator({
            rotateStringArray: true,
        })]
}
```


* ### middleware/dashboard/websocket_dashboard.js

#### Remover os Logs

### Templates
- Remover o link do `TailwindCSS`

# Deploy

## Command Build

```bash
cd code/frontend && npm i -y --omit=dev && npx webpack --mode production && npx tailwindcss -i ./tailwind.config.css -o ./static/css/tailwind.css --minify && cd ../backend && go build -o main .
```

## Start Command

```bash
cd code/backend && ./main
```