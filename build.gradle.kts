/*
 * This file was generated by the Gradle 'init' task.
 *
 * This generated file contains a sample Java project to get you started.
 * For more details take a look at the Java Quickstart chapter in the Gradle
 * User Manual available at https://docs.gradle.org/6.6.1/userguide/tutorial_java_projects.html
 */

import io.qameta.allure.gradle.AllureExtension
import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
  // Apply the java plugin to add support for Java
  java

  // Apply the java-library plugin for API and implementation separation.
  `java-library`

  // Apply the application plugin to add support for building a CLI application.
  application

  // Apply the org.jetbrains.kotlin.jvm Plugin to add support for Kotlin.
  // kotlin("multiplatform") version "1.5.0"
  kotlin("jvm") version "1.5.31"

  // Apply the scala Plugin to add support for Scala.
  scala

  // Apply the cpp-library plugin to add support for building C++ libraries
  // `cpp-library`

  // Apply the cpp-unit-test plugin to add support for building and running C++ test executables
  // `cpp-unit-test`

  // Apply the groovy Plugin to add support for Groovy.
  groovy

  // Apply the swift-library plugin to add support for building Swift libraries
  // `swift-library`

  // Apply the xctest plugin to add support for building and running Swift test executables (Linux) or bundles (macOS)
  // xctest

  // id("org.web3j") version("4.8.7")
  id("io.freefair.lombok") version ("6.2.0")

  // id("org.web3j") version ("4.8.7")

  id("io.qameta.allure") version "2.8.1"

  jacoco
  checkstyle
  pmd
  id("org.jlleitschuh.gradle.ktlint") version ("10.0.0")
  id("cz.alenkacz.gradle.scalafmt") version ("1.16.2")
  // id("com.github.sherter.google-java-format") version("0.9")
  id("com.diffplug.spotless") version "5.16.0"
}

/*
library {
  // Set the target operating system and architecture for this library
  targetMachines.add(machines.windows.x86_64)

  // Swift tool chain does not support Windows. The following targets macOS and Linux:
  targetMachines.add(machines.macOS.x86_64)

  targetMachines.add(machines.linux.x86_64)
}
*/

// TODO: check which one is legacy
java {
  toolchain {
    languageVersion.set(JavaLanguageVersion.of(Version.JAVA.id.toInt()))
  }
}
tasks.compileJava {
  options.release.set(Version.JAVA.id.toInt())
}
tasks.compileScala {
  options.release.set(Version.JAVA_FOR_SCALA.id.toInt())
}
tasks.compileTestScala {
  options.release.set(Version.JAVA_FOR_SCALA.id.toInt())
}

val compileKotlin: KotlinCompile by tasks
compileKotlin.kotlinOptions.suppressWarnings = true
tasks.withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile>().configureEach {
  kotlinOptions {
    jvmTarget = JavaVersion.VERSION_16.toString()
    suppressWarnings = true
  }
}

val compileTestKotlin: KotlinCompile by tasks
compileTestKotlin.kotlinOptions.suppressWarnings = true
tasks.withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile>().configureEach {
  kotlinOptions {
    jvmTarget = JavaVersion.VERSION_16.toString()
    suppressWarnings = true
  }
}

sourceSets {
  main {
    withConvention(ScalaSourceSet::class) {
      scala {
        setSrcDirs(listOf(SourceSet.MAIN_JAVA.path, SourceSet.MAIN_SCALA.path))
      }
    }
    java.srcDirs(
      listOf(
        SourceSet.MAIN_JAVA.path,
        SourceSet.MAIN_KOTLIN.path,
        SourceSet.MAIN_SCALA.path,
        SourceSet.MAIN_GROOVY.path
      )
    )
  }
  test {
    withConvention(ScalaSourceSet::class) {
      scala {
        setSrcDirs(listOf(SourceSet.TEST_JAVA.path, SourceSet.TEST_SCALA.path))
      }
    }
    java.srcDirs(
      listOf(
        SourceSet.TEST_JAVA.path,
        SourceSet.TEST_KOTLIN.path,
        SourceSet.TEST_SCALA.path,
        SourceSet.TEST_GROOVY.path
      )
    )
  }
}

tasks.named<AbstractCompile>("compileScala") {
  // Scala only needs the declared dependencies
  // (and not longer the output of compileJava)
  classpath = sourceSets.main.get().compileClasspath
}
tasks.named<AbstractCompile>("compileJava") {
  // Java also depends on the result of Scala compilation
  // (which automatically makes it depend of compileGroovy)
  classpath += files(sourceSets.main.get().withConvention(ScalaSourceSet::class) { scala }.classesDirectory)
}

val compileScala = tasks.named<AbstractCompile>("compileScala")
compileKotlin.dependsOn(compileScala)
compileKotlin.classpath += files(compileScala.get().destinationDir)
tasks.compileJava.get().dependsOn(compileKotlin)
tasks.named<org.gradle.jvm.tasks.Jar>("jar") {
  duplicatesStrategy = DuplicatesStrategy.EXCLUDE
}

repositories {
  // Use jcenter for resolving dependencies.
  // You can declare any Maven/Ivy/file repository here.
  // jcenter()
  mavenCentral()
  maven(url = "https://jitpack.io")
  /*maven {
      url = uri("https://mvnrepository.com")
  }*/
}

dependencies {
  // https://mvnrepository.com/artifact/org.jetbrains.kotlin/kotlin-stdlib
  implementation("org.jetbrains.kotlin:kotlin-stdlib:${Version.KOTLIN.id}")

  // Align versions of all Kotlin components
  implementation(platform("org.jetbrains.kotlin:kotlin-bom"))

  // Use the Kotlin JDK 8 standard library.
  implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")

  implementation("org.clojure:clojure:1.10.2-alpha4")

  // Use Scala 2.13 in our library project
  implementation("org.scala-lang:scala-library:${Version.SCALA.id}")

  // This dependency is exported to consumers, that is to say found on their compile classpath.
  api("org.apache.commons:commons-math3:3.6.1")

  // This dependency is used by the application.
  // implementation("com.google.guava:guava:29.0-jre")
  // This dependency is used internally, and not exposed to consumers on their own compile classpath.
  implementation("com.google.guava:guava:29.0-jre")

  // Use the latest Groovy version for building this library
  implementation("org.codehaus.groovy:groovy-all:2.5.12")

  // Use the Kotlin test library.
  // testImplementation("org.jetbrains.kotlin:kotlin-test:${Version.KOTLIN.id}")

  // Use the Kotlin JUnit integration.
  // testImplementation("org.jetbrains.kotlin:kotlin-test-junit:${Version.KOTLIN.id}")

  testImplementation("org.junit.jupiter:junit-jupiter-api:${Version.JUNIT_JUPITER.id}")
  testImplementation("org.junit.jupiter:junit-jupiter-engine:${Version.JUNIT_JUPITER.id}")
  testImplementation("org.junit.jupiter:junit-jupiter-params:${Version.JUNIT_JUPITER.id}")
  testImplementation("org.junit.vintage:junit-vintage-engine:${Version.JUNIT_JUPITER.id}")
  testImplementation("org.junit.platform:junit-platform-launcher:${Version.JUNIT_PLATFORM.id}")
  testImplementation("org.junit.platform:junit-platform-runner:${Version.JUNIT_PLATFORM.id}")
  testImplementation("org.junit.platform:junit-platform-engine:${Version.JUNIT_PLATFORM.id}")
  testImplementation("org.junit.platform:junit-platform-suite-api:${Version.JUNIT_PLATFORM.id}")

  implementation("io.cucumber:cucumber-core:${Version.CUCUMBER.id}")
  implementation("io.cucumber:cucumber-java:${Version.CUCUMBER.id}")
  implementation("io.cucumber:cucumber-gherkin:${Version.CUCUMBER.id}")
  implementation("io.cucumber:gherkin:18.1.0")

  testImplementation("io.cucumber:cucumber-junit:${Version.CUCUMBER.id}")
  implementation("io.cucumber:cucumber-junit-platform-engine:${Version.CUCUMBER.id}")
  implementation("io.cucumber:cucumber-java8:${Version.CUCUMBER.id}")

  // Use the awesome Spock testing and specification framework even with Java
  testImplementation("org.spockframework:spock-core:1.3-groovy-2.5")
  testImplementation("junit:junit:${Version.JUNIT4.id}")

  // Use Scalatest for testing our library
  // testImplementation("junit:junit:4.12")
  testImplementation("org.scalatest:scalatest_2.13:${Version.SCALA_TEST.id}")
  testImplementation("org.scalatestplus:junit-4-12_2.13:${Version.SCALA_TEST_PLUS.id}")
  testImplementation("org.scalatest:scalatest-freespec_2.13:${Version.SCALA_TEST.id}")
  testImplementation("org.scalatest:scalatest-funsuite_2.13:${Version.SCALA_TEST.id}")

  // Need scala-xml at test runtime
  testRuntimeOnly("org.scala-lang.modules:scala-xml_2.13:1.2.0")

  // testImplementation("io.qameta.allure:allure-java-commons:${Version.ALLURE.id}")
  implementation("io.qameta.allure:allure-java-commons:${Version.ALLURE.id}")
  implementation("io.qameta.allure:allure-cucumber6-jvm:${Version.ALLURE.id}")
  testImplementation("io.qameta.allure:allure-junit4:${Version.ALLURE.id}")

  implementation("io.qameta.allure:allure-attachments:${Version.ALLURE.id}")

  runtimeOnly("com.pinterest.ktlint:ktlint-core:${Version.KTLINT.id}")
  runtimeOnly("com.pinterest.ktlint:ktlint-ruleset-standard:${Version.KTLINT.id}")
  runtimeOnly("com.pinterest.ktlint:ktlint-reporter-plain:${Version.KTLINT.id}")

  // implementation("cz.alenkacz.gradle.scalafmt:cz.alenkacz.gradle.scalafmt.gradle.plugin:${Version.SCALA_FMT.id}")

  compileOnly("org.projectlombok:lombok:${Version.LOMBOK.id}")
  annotationProcessor("org.projectlombok:lombok:${Version.LOMBOK.id}")

  testCompileOnly("org.projectlombok:lombok:${Version.LOMBOK.id}")
  testAnnotationProcessor("org.projectlombok:lombok:${Version.LOMBOK.id}")

  implementation("io.github.cdimascio:java-dotenv:5.2.2")

  implementation("org.web3j:core:${Version.WEB3J.id}")
  implementation("org.web3j:crypto:${Version.WEB3J.id}")
  implementation("org.web3j:utils:${Version.WEB3J.id}")
  implementation("org.web3j:abi:${Version.WEB3J.id}")
  implementation("org.web3j:codegen:${Version.WEB3J.id}")

  runtimeOnly("com.github.komputing.kethereum:crypto_api:${Version.KETHEREUM.id}")
  runtimeOnly("com.github.komputing.kethereum:model:${Version.KETHEREUM.id}")
  runtimeOnly("com.github.komputing.kethereum:extensions_kotlin:${Version.KETHEREUM.id}")
}

// TODO: needed for Kotlin and kotlin.test 1.5.0
/*
val testCompile: Configuration by configurations.creating
configurations {
  testCompile.extendsFrom(testImplementation.get())
}
*/

configure<AllureExtension> {
  autoconfigure = true
  aspectjweaver = true
  version = Version.ALLURE.id
  allureJavaVersion = Version.JAVA.id
  configuration = "testImplementation"

  clean = true

  resultsDir = file("../../report/allure-results")
  reportDir = file("../../report/allure-reports")

  useJUnit5 {
    version = Version.ALLURE.id
  }
}

tasks.test {
  filter {
    exclude("/e2e/**")
    exclude("TestRunner")
    useJUnitPlatform()
    // TODO: remove this one if not needed
    testLogging.showStandardStreams = true
  }
  maxParallelForks = 3
/*
val test by tasks.getting(Test::class) {
  ignoreFailures = true
  // Use junit platform for unit tests
  useJUnitPlatform()
  testLogging.showStandardStreams = true */
  systemProperty("junit.jupiter.execution.parallel.enabled", "true")
  systemProperty("junit.jupiter.execution.parallel.config.strategy", "dynamic")
  systemProperty("junit.jupiter.extensions.autodetection.enabled", "true")
  // systemProperty("allure.results.directory", "../../allure-results")
}
checkstyle {
  toolVersion = "9.0"
  // ignoreFailures = false
  maxWarnings = 0
}
tasks.withType<Checkstyle>().configureEach {
  source("src/special_dir")
  include("**/*.java")
  exclude("**/*Template.java")
  exclude("src/test/template_*")
  exclude("node_modules")
  reports {
    xml.isEnabled = false
    html.isEnabled = true
    // html.stylesheet = resources.text.fromFile("config/xsl/checkstyle-custom.xsl")
  }
}

// tasks.googleJavaFormat {
//   // source = sourceSets*.allJava
//   // source = sourceSets.get(allJava)
//   source("src/special_dir")
//   include("**/*.java")
//   exclude("**/*Template.java")
//   exclude("src/test/template_*")
//   exclude("node_modules/**/*.java")
// }

tasks.withType<Pmd>() {
  isConsoleOutput = true
  // toolVersion = Version.PMD.id
  rulesMinimumPriority.set(5)
  ruleSets = listOf("category/java/errorprone.xml", "category/java/bestpractices.xml")
}

scalafmt {
  // .scalafmt.conf in the project root is default value, provide only if other location is needed
  // config file has to be relative path from current project or root project in case of multimodule projects
  // example usage: 
  // configFilePath = ".scalafmt.conf"
}

configure<com.diffplug.gradle.spotless.SpotlessExtension> {
  // optional: limit format enforcement to just the files changed by this feature branch
  // ratchetFrom("origin/main")

  format(
    "misc",
    {
      // define the files to apply `misc` to
      target("*.gradle", "*.md", ".gitignore")

      // define the steps to apply to those files
      trimTrailingWhitespace()
      indentWithTabs() // or spaces. Takes an integer argument if you don't like 4
      endWithNewline()
    }
  )

  // java {
  //   // don't need to set target, it is inferred from java

  //   // apply a specific flavor of google-java-format
  //   googleJavaFormat("1.11.0").aosp().reflowLongStrings()
  //   // make sure every file has the following copyright header.
  //   // optionally, Spotless can set copyright years by digging
  //   // through git history (see "license" section below)
  //   // licenseHeader("/* (C)$YEAR */")
  // }
}

tasks.named<Wrapper>("wrapper") {
  gradleVersion = Version.GRADLE.id
  distributionType = Wrapper.DistributionType.ALL
}

enum class Version(val id: String) {
  GATLING("3.6.1"),
  JUNIT_JUPITER("5.8.1"),
  JUNIT_PLATFORM("1.8.1"),
  JUNIT4("4.13.2"),
  SCALA("2.13.6"),
  SCALA_TEST("3.2.10"),
  SCALA_TEST_PLUS("3.2.0.0"),
  JACKSON("2.12.5"),
  SNAKEYAML("1.29"),
  JOOQ("3.15.3"),
  POSTGRESQL("42.2.24"),
  REST_ASSURED("4.4.0"),
  HAMCREST("2.2"),
  JAVAFAKER("1.0.2"),
  AWAITILITY("4.1.0"),
  KETHEREUM("0.83.4"),
  CUCUMBER("6.11.0"),
  CUCUMBER_JUNIT("6.11.0"),
  ALLURE("2.15.0"),
  ALLURE_GRADLE("2.8.1"),
  JAVA("16"),
  JAVA_FOR_SCALA("11"),
  KOTLIN("1.5.31"),
  LOMBOK("1.18.20"),
  WEB3J("5.0.0"),
  GRADLE("7.2"),
  PMD("6.39.0"),
  KTLINT_GRADLE_PLUGIN("10.0.0"),
  KTLINT("0.42.1"),
  SCALA_FMT("1.16.2");
}

enum class SourceSet(val path: String) {
  MAIN_JAVA("src/main/java"),
  MAIN_KOTLIN("src/main/kotlin"),
  MAIN_SCALA("src/main/scala"),
  MAIN_GROOVY("src/main/groovy"),
  TEST_JAVA("src/test/java"),
  TEST_KOTLIN("src/test/kotlin"),
  TEST_SCALA("src/test/scala"),
  TEST_GROOVY("src/test/groovy");
}
