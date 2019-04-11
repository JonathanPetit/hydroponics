import React, { Component } from 'react';
import axios from 'axios';
import {Line} from 'react-chartjs-2';

class Humidity extends Component {
    constructor(props) {
        super(props);
        this.state = {
          humidities: []
        };
    }

    componentDidMount() {
        axios.get("http://192.168.0.13:8080/humidities")
        .then(res => {
            const hums = res.data;
            this.setState({humidities: hums});})
        .catch(error => console.log("Error Temp"));
    }

    render() {
        const humidity_chart = {
            labels: this.state.humidities.map(hum => Date(hum.savedAt)),
            datasets: [
              {
                label: 'Humidity',
                fill: false,
                lineTension: 0.1,
                backgroundColor: 'rgba(75,192,192,0.4)',
                borderColor: 'rgba(75,192,192,1)',
                borderCapStyle: 'butt',
                borderDash: [],
                borderDashOffset: 0.0,
                borderJoinStyle: 'miter',
                pointBorderColor: 'rgba(75,192,192,1)',
                pointBackgroundColor: '#fff',
                pointBorderWidth: 1,
                pointHoverRadius: 5,
                pointHoverBackgroundColor: 'rgba(75,192,192,1)',
                pointHoverBorderColor: 'rgba(220,220,220,1)',
                pointHoverBorderWidth: 2,
                pointRadius: 1,
                pointHitRadius: 10,
                data: this.state.humidities.map(hum => hum.value)
              }
            ]
          };
          
        return (
            <div><Line data={humidity_chart}/></div>
        )
    }
}

export default Humidity;