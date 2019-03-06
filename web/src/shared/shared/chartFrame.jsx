import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Chart from 'chart.js';

import './chartFrame.css';

export class ChartFrame extends Component {
  state = {
    showButtons: false,
    isChart: false,
  }

  chartData = null;

  static propTypes = {
    object: PropTypes.objectOf(PropTypes.object).isRequired,
    children: PropTypes.node.isRequired,
  }

  handleMouseEnter = (e) => {
    e.stopPropagation();
    this.setState({ showButtons: true });
  }

  handleMouseLeave = (e) => {
    e.stopPropagation();
    this.setState({ showButtons: false });
  }

  handleChangeView = () => {
    this.setState(prev => ({
      isChart: !prev.isChart,
    }));
    this.renderChart();
  }

  renderChart() {
    const ctx = document.getElementById('chart');
    const numberData = [];
    const labels = [];
    Object.keys(this.props.object).forEach((key) => {
      numberData.push(this.props.object[key]);
      labels.push(key);
    });
    this.chartData = {
      datasets: [{ data: numberData }],
      labels,
    };
    const chart = new Chart(ctx, {
      type: 'pie',
      data: this.chartData,
    });
  }

  render() {
    return (
      <div className="container" onMouseEnter={this.handleMouseEnter} onMouseLeave={this.handleMouseLeave}>
        {this.state.showButtons
          ? (
            <div className="buttons">
              <button type="button" onClick={this.handleChangeView}>Pie</button>
            </div>
          )
          : null
        }
      </div>
    );
  }
}
