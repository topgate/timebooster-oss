// groovy 2.9
// groovy ./script/swagger-codegen.groovy
@GrabResolver(name = "eaglesakura", root = "https://dl.bintray.com/eaglesakura/maven/")
@Grab("com.eaglesakura:lightweight-swagger-codegen:1.0.52")
import java.lang.Object;

static main(String[] args) {
    def BUILD_PATH = "./api/build"

    new File("${BUILD_PATH}/").mkdirs()
    new File("${BUILD_PATH}/swagger.yaml").text = new File("api/swagger.yaml").text
            .replaceAll("__GENERATED_DATE__", "${new Date()}")

    // Generate swagger.json
    com.eaglesakura.swagger.generator.Generator.main([
            "generate",
            "-l", "io.swagger.codegen.languages.SwaggerGenerator",
            "-o", "${BUILD_PATH}",
            "-i", "${BUILD_PATH}/swagger.yaml",
    ] as String[])

    // Generate GAE/Go server binding
    com.eaglesakura.swagger.generator.Generator.main([
            "generate",
            "-l", "com.eaglesakura.swagger.generator.GoServerCodegen",
            "-o", "${BUILD_PATH}/server",
            "-i", "${BUILD_PATH}/swagger.yaml",
            "-c", "api/config.json",
    ] as String[])

    // Generate GAE/Go client binding
    com.eaglesakura.swagger.generator.Generator.main([
            "generate",
            "-l", "com.eaglesakura.swagger.generator.GoClientCodegen",
            "-o", "${BUILD_PATH}/client",
            "-i", "${BUILD_PATH}/swagger.yaml",
            "-c", "api/config.json",
    ] as String[])

    // formatting
    // Go Formatting   : go fmt
    println(["go", "fmt", "${BUILD_PATH}/server/..."].execute().text)
    println(["go", "fmt", "${BUILD_PATH}/client/..."].execute().text)
}
