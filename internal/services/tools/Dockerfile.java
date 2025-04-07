FROM eclipse-temurin:17-jdk-jammy

# Устанавливаем инструменты
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl unzip \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем Checkstyle с полной конфигурацией
RUN curl -L https://github.com/checkstyle/checkstyle/releases/download/checkstyle-10.12.4/checkstyle-10.12.4-all.jar \
    -o /opt/checkstyle.jar

# Устанавливаем PMD
RUN curl -L https://github.com/pmd/pmd/releases/download/pmd_releases%2F6.55.0/pmd-bin-6.55.0.zip \
    -o /opt/pmd.zip && \
    unzip /opt/pmd.zip -d /opt && \
    mv /opt/pmd-bin-6.55.0 /opt/pmd && \
    rm /opt/pmd.zip

# Полная конфигурация Checkstyle
COPY <<EOF /opt/checkstyle.xml
<?xml version="1.0"?>
<!DOCTYPE module PUBLIC "-//Checkstyle//DTD Checkstyle Configuration 1.3//EN"
    "https://checkstyle.org/dtds/configuration_1_3.dtd">
<module name="Checker">
  <module name="TreeWalker">
    <module name="AvoidStarImport"/>
    <module name="UnusedImports"/>
    <module name="UnusedLocalVariable"/>
    <module name="ConstantName"/>
    <module name="EmptyBlock"/>
    <module name="MagicNumber"/>
    <module name="StringConcatenation"/>
    <module name="EqualsAvoidNull"/>
    <module name="DuplicateCode"/>
  </module>
</module>
EOF

# Скрипт запуска линтеров
COPY <<EOF /usr/local/bin/lint-java
#!/bin/sh
echo "Running Checkstyle..."
java -jar /opt/checkstyle.jar -c /opt/checkstyle.xml /code/Main.java

echo "Running PMD..."
/opt/pmd/bin/run.sh pmd -d /code -R rulesets/java/quickstart.xml
EOF

RUN chmod +x /usr/local/bin/lint-java

WORKDIR /code
ENTRYPOINT ["lint-java"]