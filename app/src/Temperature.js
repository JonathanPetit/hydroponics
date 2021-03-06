import React, { Component } from 'react';
import axios from 'axios';
import {Line} from 'react-chartjs-2';

class Temperature extends Component {
    constructor(props) {
        super(props);
        this.state = {
          temperatures: []
        };
    }

    componentDidMount() {
        axios.get("http://192.168.0.13:8080/temperatures")
        .then(res => {
            const temps = res.data;
            this.setState({temperatures: temps});})
        .catch(error => console.log("Error Temp"));
    }

    render() {
        const temperature_chart = {
            labels: this.state.temperatures.map(temp => Date(temp.savedAt)),
            datasets: [
              {
                label: 'Temperature',
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
                data: this.state.temperatures.map(temp => temp.value)
              }
            ]
          };
          
        return (
            <div><Line data={temperature_chart}/></div>
        )
    }
}

export default Temperature;