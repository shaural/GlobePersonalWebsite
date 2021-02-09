<template>
  <div id="chartContainer" ref="chartdiv"></div>
</template>

<script>
import * as am4core from "@amcharts/amcharts4/core";
import * as am4maps from "@amcharts/amcharts4/maps";
import am4geodata_worldLow from "@amcharts/amcharts4-geodata/worldLow";
import am4geodata_usaLow from "@amcharts/amcharts4-geodata/usaLow";
import am4geodata_indiaLow from "@amcharts/amcharts4-geodata/indiaLow";

const visited_countries = [
  "BE",
  "FR",
  "DE",
  "NL",
  "IT",
  "GB",
  "CA",
  "SG",
  "AT",
  "ES",
  "PL",
  "CH",
  "EG",
  "GR",
  "DO",
  "MA",
  "TR",
  "PT",
  "LU",
  "MC"
];
const visited_states_usa = [
  "US-CA",
  "US-TX",
  "US-FL",
  "US-AK",
  "US-WA",
  "US-NV",
  "US-MA",
  "US-OH",
  "US-WI",
  "US-CO",
  "US-IN",
  "US-IL",
  "US-NJ",
  "US-MN",
  "US-VA",
  "US-NY",
  "US-MO"
];
const visited_states_india = [
  "IN-GJ",
  "IN-MH",
  "IN-GA",
  "IN-DL",
  "IN-UP",
  "IN-PB",
  "IN-RJ",
  "IN-KA"
];

export default {
  name: "Globe",
  mounted() {
    // Create map instance
    var chart = am4core.create(this.$refs.chartdiv, am4maps.MapChart);
    setupChart(chart);
    plotBgAndLines(chart);
    plotCountries(chart);
    plotStatesUSA(chart);
    plotStatesIndia(chart);
    // animate(chart);
  }
};

function setupChart(chart) {
  chart.width = am4core.percent(100);
  chart.height = am4core.percent(100);

  chart.panBehavior = "rotateLongLat";

  // limits vertical rotation
  chart.adapter.add("deltaLatitude", function(delatLatitude) {
    return am4core.math.fitToRange(delatLatitude, -90, 90);
  });

  // Set map definition
  chart.geodata = am4geodata_worldLow;

  // Set projection
  chart.projection = new am4maps.projections.Orthographic();
}

function plotBgAndLines(chart) {
  // latitude and longitude lines
  var graticuleSeries = chart.series.push(new am4maps.GraticuleSeries());
  graticuleSeries.mapLines.template.line.stroke = am4core.color("#67b7dc");
  graticuleSeries.mapLines.template.line.strokeOpacity = 0.2;
  graticuleSeries.fitExtent = false;

  // Background -> water
  chart.backgroundSeries.mapPolygons.template.polygon.fill = am4core.color(
    "#aadaff"
  );
  chart.backgroundSeries.mapPolygons.template.polygon.fillOpacity = 1;
}

function plotCountries(chart) {
  // Not Visited (no hover)
  // Create map polygon series
  var polygonSeriesDisbaled = chart.series.push(new am4maps.MapPolygonSeries());
  polygonSeriesDisbaled.north = 90;

  // Make map load polygon (like country names) data from GeoJSON
  polygonSeriesDisbaled.useGeodata = true;

  // Display only needed countries
  polygonSeriesDisbaled.exclude = visited_countries;

  // Configure series
  var polygonTemplateDisabled = polygonSeriesDisbaled.mapPolygons.template;
  // polygonTemplateDisabled.tooltipText = "{name}";
  polygonTemplateDisabled.fill = am4core.color("#404F24"); // old color: #74B266
  polygonTemplateDisabled.opacity = 0.2;

  // Visited (with hover)
  // Create map polygon series
  var polygonSeries = chart.series.push(new am4maps.MapPolygonSeries());
  polygonSeries.north = 90;

  // Make map load polygon (like country names) data from GeoJSON
  polygonSeries.useGeodata = true;

  // Display only needed countries
  polygonSeries.include = visited_countries;

  // Configure series
  var polygonTemplate = polygonSeries.mapPolygons.template;
  polygonTemplate.tooltipText = "{name}";
  polygonTemplate.fill = am4core.color("#74B266"); // old color: #74B266
  polygonTemplate.stroke = am4core.color("#000000"); // outline color
  polygonTemplate.strokeWidth = 0.5;
  polygonTemplate.strokeOpacity = 0.8;

  // Create hover state and set alternative fill color
  var hs = polygonTemplate.states.create("hover");
  hs.properties.fill = am4core.color("#367B25");

  // chart.deltaLongitude = -80;
  // // Zoom
  // // NOTE: this will only work when the polygon is included in the series (USA and INDIA -> removed from global series)
  // chart.events.on("ready", function() {
  //   var india = polygonSeries.getPolygonById("IN");

  //   // Pre-zoom
  //   chart.zoomToMapObject(india);

  //   // Set active state
  //   // setTimeout(function() {
  //   //   india.isActive = true;
  //   // }, 1000);
  // });
}

function plotStatesUSA(chart) {
  let usaSeries = chart.series.push(new am4maps.MapPolygonSeries());
  usaSeries.geodata = am4geodata_usaLow;

  // Display only needed states
  usaSeries.include = visited_states_usa;

  let usPolygonTemplate = usaSeries.mapPolygons.template;
  usPolygonTemplate.tooltipText = "{name}";
  // usPolygonTemplate.fill = chart.colors.getIndex(1);
  // usPolygonTemplate.nonScalingStroke = true;
  usPolygonTemplate.fill = am4core.color("#74B266"); // old color: #74B266
  usPolygonTemplate.stroke = am4core.color("#000000"); // outline color
  usPolygonTemplate.strokeWidth = 0.5;
  usPolygonTemplate.strokeOpacity = 0.8;

  // Hover state
  let hs = usPolygonTemplate.states.create("hover");
  hs.properties.fill = am4core.color("#367B25");
}

function plotStatesIndia(chart) {
  let indiaSeries = chart.series.push(new am4maps.MapPolygonSeries());
  indiaSeries.geodata = am4geodata_indiaLow;

  // Display only needed states
  indiaSeries.include = visited_states_india;

  let indiaPolygonTemplate = indiaSeries.mapPolygons.template;
  indiaPolygonTemplate.tooltipText = "{name}";
  indiaPolygonTemplate.fill = am4core.color("#74B266"); // old color: #74B266
  indiaPolygonTemplate.stroke = am4core.color("#000000"); // outline color
  indiaPolygonTemplate.strokeWidth = 0.5;
  indiaPolygonTemplate.strokeOpacity = 0.8;

  // Hover state
  let hs = indiaPolygonTemplate.states.create("hover");
  hs.properties.fill = am4core.color("#367B25");
}

// function animate(chart) {
//   let animation;
//   setTimeout(function() {
//     animation = chart.animate(
//       { property: "deltaLongitude", to: 100000 },
//       20000000
//     );
//   }, 3000);

//   chart.seriesContainer.events.on("down", function() {
//     if (animation) {
//       animation.stop();
//     }
//   });

//   // chart.seriesContainer.events.on("up", function() {
//   //   animation = chart.animate(
//   //     { property: "deltaLongitude", to: 100000 },
//   //     20000000
//   //   );
//   // });
// }
</script>

<style scoped>
#chartContainer {
  width: 75%;
  height: 75%;
  position: absolute;
  margin: auto;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  overflow: hidden;
  background: #000;
}
</style>
