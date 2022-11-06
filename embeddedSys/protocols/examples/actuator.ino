int relayPin = 12;//define a pin for relay
String buttonTitle1 ="Light ON";
String buttonTitle2 ="Light OFF";

int relayStat = 1;//initial state . 1 ON, 0 OFF

#include <WiFi.h>
#include <WiFiClient.h>
#include <WebServer.h>
#include <ESPmDNS.h>

const char *ssid = "YOUR_WIFI_SSID";
const char *password = "YOUR_WIFI_PASSWORD";

WebServer server(80);
const int led = 13;

void handleRoot() {
  //Robojax.com ESP32 Relay Control 
  digitalWrite(led, 1);  
 String HTML ="<!DOCTYPE html>\
  <html>\
  <head>\
  	
<title>Robojax ESP32 Relay Control</title>\
  	
<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\
  
<style>\
 
html,body{	
width:100%\;
height:100%\;
margin:0}
*{box-sizing:border-box}
.colorAll{
	background-color:#90ee90}
.colorBtn{
	background-color:#add8e6}
.angleButtdon,a{
	font-size:30px\;
border:1px solid #ccc\;
display:table-caption\;
padding:7px 10px\;
text-decoration:none\;
cursor:pointer\;
padding:5px 6px 7px 10px}a{
	display:block}
.btn{
	margin:5px\;
border:none\;
display:inline-block\;
vertical-align:middle\;
text-align:center\;
white-space:nowrap}
";
   
  HTML +="</style>

</head>

<body>
<h1>Robojax ESP32 Relay Control </h1>
";
  
  if(relayStat){
    HTML +="	<div class=\"btn\">
		<a class=\"angleButton\" style=\"background-color:#90ee90\"  href=\"/relay?do=off\">";
    HTML +=buttonTitle1; //Light ON title
  }else{
    HTML +="	<div class=\"btn\">
		<a class=\"angleButton \" style=\"background-color:#f56464\"  href=\"/relay?do=on\">";
    HTML +=buttonTitle2;//Light OFF title  
  }
  HTML +="</a>	
	</div>
</body>
</html>
";
  server.send(200, "text/html", HTML);  
  digitalWrite(led, 0);  
  //Robojax.com ESP32 Relay Control 
}//handleRoot()

void handleNotFound() {
  //Robojax.com ESP32 Relay Control 
  digitalWrite(led, 1);
  String message = "File Not Found

";
  message += "URI: ";
  message += server.uri();
  message += "
Method: ";
  message += (server.method() == HTTP_GET) ? "GET" : "POST";
  message += "
Arguments: ";
  message += server.args();
  message += "
";

  for (uint8_t i = 0; i < server.args(); i++) {
    message += " " + server.argName(i) + ": " + server.arg(i) + "
";
  }

  server.send(404, "text/plain", message);
  digitalWrite(led, 0);
  //Robojax.com ESP32 Relay Control 
}

void setup(void) {
  //Robojax.com ESP32 Relay Control 
  pinMode(led, OUTPUT);
  digitalWrite(led, 0);  
  pinMode(relayPin, OUTPUT);// define a pin as output for relay
  digitalWrite(relayPin, relayStat);//initial state either ON or OFF
  Serial.begin(115200);//initialize the serial monitor
  Serial.println("Robojax ESP32 Relay");

  //Robojax.com ESP32 Relay Control 

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  Serial.println("");
 // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  //multicast DNS //Robojax.com ESP32 Relay Control 
  if (MDNS.begin("robojaxEsp32")) {
    Serial.println("MDNS responder started");
  }
  server.on("/", handleRoot);
  server.on("/relay", HTTP_GET, relayControl);  //  relayControl() is at the end of this code     
  server.onNotFound(handleNotFound);
  server.begin();
  Serial.println("HTTP server started");

}

void loop(void) {
  //Robojax ESP32 Relay Control . 
  server.handleClient();

  if(relayStat)
  {
    digitalWrite(relayPin, LOW);
  }else{
    digitalWrite(relayPin, HIGH);    
  }

  //Robojax ESP32 Relay Control 
  delay(100);
}

/*
 * relayControl()
 * updates the value of "relayStat" varible to 1 or zero 
 * returns nothing
 * written by Ahmad Shamshiri
 * on Wednesday Feb 19, 2020 at 20:34 in Ajax, Ontario, Canada
 * www.robojax.com
 */
void relayControl() {
  if(server.arg("do") == "on")
  {
    relayStat =1;
    

  }else{
    relayStat =0;   

  }

  handleRoot();
}//relayControl() end