<template>
  <div id="chartContainer" ref="chartdiv"></div>
</template>

<script>
import * as am4core from "@amcharts/amcharts4/core";
import * as am4maps from "@amcharts/amcharts4/maps";
import am4geodata_worldLow from "@amcharts/amcharts4-geodata/worldLow";
const visited_countries = [
  "FR",
  "DE",
  "BE",
  "NL",
  "IT",
  "GB",
  "US",
  "CA",
  "IN",
  "SG",
  "AT"
];

export default {
  name: "Globe",
  mounted() {
    // Create map instance
    var chart = am4core.create(this.$refs.chartdiv, am4maps.MapChart);

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

    // Add countries
    plotCountries(chart);

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

    let animation;
    setTimeout(function() {
      animation = chart.animate(
        { property: "deltaLongitude", to: 100000 },
        20000000
      );
    }, 3000);

    chart.seriesContainer.events.on("down", function() {
      if (animation) {
        animation.stop();
      }
    });

    // chart.seriesContainer.events.on("up", function() {
    //   animation = chart.animate(
    //     { property: "deltaLongitude", to: 100000 },
    //     20000000
    //   );
    // });
  }
};
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
}
</script>

<style scoped>
#chartContainer {
  width: 99%;
  height: 96vh;
  position: relative;
  overflow: hidden;
  padding: 10px;
}
</style>
