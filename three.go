//Struct(User defined types)
//A car game 


/*2 type of methods (A Member function)
1. Value receiver (Just takes the value)
	Why would we use value receiver ?
		It makes a copy. It never interacts with original copy.
2. Pointer receiver (To change the values inside the struct)
*/
package main 
import ("fmt")

const usixteenbitmax float64=65535
const kmh_mutiple float64=1.60934    //1mile = 1.60934 km

type car struct{
	gas_pedal uint16	//0 to 2^16-1
	brake_pedal uint16
	steering_wheel int16 //-32K to +32K
	top_speed_kmh float64
}
//we are associating this func kmh() to struct car which make it a method
//value reciver method
func (c car) kmph() float64 {
	return float64(c.gas_pedal)*(c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal)*(c.top_speed_kmh/usixteenbitmax/kmh_mutiple)
}

//pointer reciver method
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh=newspeed
}
//A Separate fucntion does exact same thing as above method
func newer_top_speed(c car, speed float64) car{
	c.top_speed_kmh=speed;
	return c;
}

func main() {
	a_car :=car{gas_pedal:22341, 
				brake_pedal:0, 
				steering_wheel:12456, 
				top_speed_kmh:225.0 }

	b_car :=car{22341, 0, 12456, 225.0}

	// fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmph())
	fmt.Println(b_car.mph())
	// fmt.Println(b_car.brake_pedal)
	a_car = newer_top_speed(a_car, 600)
	b_car.new_top_speed(500)
	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(b_car.top_speed_kmh)
}