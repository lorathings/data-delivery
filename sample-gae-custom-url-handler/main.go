//Sample golang code for appengine project to handle Sertone payloads
//Related to data delivery using a Custom URL
//https://dev.sertone.com/how-to-send-payload-to-a-custom-url/

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"appengine"
	"encoding/base64"
)

// DeliveryPayload represents the actual payloads sent to application on uplink
type DeliveryPayload struct {                                                  
 Payload  []byte                 `json:"payload"`                           
 Fields   map[string]interface{} `json:"fields,omitempty"`                  
 FPort    uint8                  `json:"port,omitempty"`                    
 FCnt     uint32                 `json:"counter"`                           
 DevEUI   string                 `json:"dev_eui"`                           
 Metadata []AppMetadata          `json:"metadata"`                          
} 

// AppMetadata represents gathered metadata that are sent to gateways
type AppMetadata struct {                                            
 Frequency  float32 `json:"frequency"`                               
 DataRate   string  `json:"datarate"`                                
 CodingRate string  `json:"codingrate"`                              
 Timestamp  uint32  `json:"gateway_timestamp"`                       
 Time       string  `json:"gateway_time,omitempty"`                  
 Channel    uint32  `json:"channel"`                                 
 ServerTime string  `json:"server_time"`                             
 Rssi       int32   `json:"rssi"`                                    
 Lsnr       float32 `json:"lsnr"`                                    
 RFChain    uint32  `json:"rfchain"`                                 
 CRCStatus  int32   `json:"crc"`                                     
 Modulation string  `json:"modulation"`                              
 GatewayEUI string  `json:"gateway_eui"`
 Altitude   int32   `json:"altitude"`   
 Longitude  float32 `json:"longitude"`  
 Latitude   float32 `json:"latitude"`   
}

func init() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/sertone", pushHandler)
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("pushHandler()")
	c.Infof("http.Request: %v", r)
	
	//Get method
	method := r.Method
	c.Infof("method: %v", method)
	c.Infof("r.Header: %v", r.Header)
	
	if method == "POST" {
	
		c.Infof("r.Body: %v", r.Body)
		
		dp := &DeliveryPayload{}
		if err := json.NewDecoder(r.Body).Decode(dp); err != nil {
			c.Infof("Could not decode body: %v", err)
			return
		}
		c.Infof("data: %v", dp)	
		
		//field1, //payload //byte
		field1 := dp.Payload
		c.Infof("dp.Payload: %v", string(field1))
		
		//field2, //port //int
		field2 := dp.FPort
		c.Infof("dp.FPort: %v", field2)
		
		//field3, //counter //int32
		field3 := dp.FCnt
		c.Infof("dp.FCnt: %v", field3)
		
		//field4, //dev_eui //string
		field4 := dp.DevEUI
		c.Infof("dp.DevEUI: %v", field4)
		
		//field5, //gateway_timestamp //int32
		field5 := dp.Metadata[0].Timestamp 
		c.Infof("dp.Timestamp: %v", field5)
		
		//field6, //gateway_time //string
		field6 := dp.Metadata[0].Time
		c.Infof("dp.Time: %v", field6)
		
		//field7, //modulation //string
		field7 := dp.Metadata[0].Modulation
		c.Infof("dp.Modulation: %v", field7)
		
		//field8, //gateway_eui //string
		field8 := dp.Metadata[0].GatewayEUI
		c.Infof("dp.GatewayEUI: %v", field8)
		
		//elevation, 	//int32
		field9 := dp.Metadata[0].Altitude
		c.Infof("dp.Altitude: %v", field9)
		
		//latitude,	//float32
		field10 := dp.Metadata[0].Latitude
		c.Infof("dp.Latitude: %v", field10)
		
		//longitude	//float32
		field11 := dp.Metadata[0].Longitude
		c.Infof("dp.Longitude: %v", field11)
	
	} else {
		
		q_data := r.FormValue("q")
		c.Infof("q_data: %v", q_data)
		
		thisData, err := base64.StdEncoding.DecodeString(q_data)
		if err != nil {
			c.Infof("err: %v", err)
		}
		c.Infof("thisData: %v", thisData)	
		
		dp := new(DeliveryPayload)
		err = json.Unmarshal(thisData, dp)
		if err != nil {
			c.Infof("err: %v", err)
		}
		//field1, //payload //byte
		field1 := dp.Payload
		c.Infof("dp.Payload: %v", string(field1))
		
		//field2, //port //int
		field2 := dp.FPort
		c.Infof("dp.FPort: %v", field2)
		
		//field3, //counter //int32
		field3 := dp.FCnt
		c.Infof("dp.FCnt: %v", field3)
		
		//field4, //dev_eui //string
		field4 := dp.DevEUI
		c.Infof("dp.DevEUI: %v", field4)
		
		//field5, //gateway_timestamp //int32
		field5 := dp.Metadata[0].Timestamp 
		c.Infof("dp.Timestamp: %v", field5)
		
		//field6, //gateway_time //string
		field6 := dp.Metadata[0].Time
		c.Infof("dp.Time: %v", field6)
		
		//field7, //modulation //string
		field7 := dp.Metadata[0].Modulation
		c.Infof("dp.Modulation: %v", field7)
		
		//field8, //gateway_eui //string
		field8 := dp.Metadata[0].GatewayEUI
		c.Infof("dp.GatewayEUI: %v", field8)
		
		//elevation, 	//int32
		field9 := dp.Metadata[0].Altitude
		c.Infof("dp.Altitude: %v", field9)
		
		//latitude,	//float32
		field10 := dp.Metadata[0].Latitude
		c.Infof("dp.Latitude: %v", field10)
		
		//longitude	//float32
		field11 := dp.Metadata[0].Longitude
		c.Infof("dp.Longitude: %v", field11)
	
	}
	
	fmt.Fprint(w, "Payload processed!")
	
}