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
const latitude_map = new Map([
  ["IN", -80],
  ["US", 100],
  ["BE", 0],
  ["FR", 0],
  ["DE", -10],
  ["NL", 0],
  ["IT", -10],
  ["GB", 0],
  ["CA", 100],
  ["SG", -100],
  ["AT", -10],
  ["ES", 0],
  ["PL", -10],
  ["CH", 0],
  ["EG", -30],
  ["GR", -10],
  ["DO", 80],
  ["MA", 0],
  ["TR", -30],
  ["PT", 0],
  ["LU", 0],
  ["MC", 0]
]);

export default {
  name: "Globe",
  data: function() {
    return {
      chart: undefined,
      animation: undefined,
      visited_countries: visited_countries,
      visited_states_usa: visited_states_usa,
      visited_states_india: visited_states_india,
      latitude_map: latitude_map
    };
  },
  methods: {
    setupChart() {
      this.chart.width = am4core.percent(100);
      this.chart.height = am4core.percent(100);

      this.chart.panBehavior = "rotateLongLat";

      // limits vertical rotation
      this.chart.adapter.add("deltaLatitude", function(delatLatitude) {
        return am4core.math.fitToRange(delatLatitude, -90, 90);
      });

      // Set map definition
      this.chart.geodata = am4geodata_worldLow;

      // Set projection
      this.chart.projection = new am4maps.projections.Orthographic();

      // Set home for chart.GoHome() (to zoom out)
      this.chart.homeZoomLevel = 0;

      // Add button to zoom out
      var home = this.chart.chartContainer.createChild(am4core.Button);
      home.label.text = "Home";
      home.align = "right";
      home.events.on("hit", () => {
        this.chart.goHome();
      });
    },
    plotBgAndLines() {
      // latitude and longitude lines
      var graticuleSeries = this.chart.series.push(
        new am4maps.GraticuleSeries()
      );
      graticuleSeries.mapLines.template.line.stroke = am4core.color("#67b7dc");
      graticuleSeries.mapLines.template.line.strokeOpacity = 0.2;
      graticuleSeries.fitExtent = false;

      // Background -> water
      this.chart.backgroundSeries.mapPolygons.template.polygon.fill = am4core.color(
        "#aadaff"
      );
      this.chart.backgroundSeries.mapPolygons.template.polygon.fillOpacity = 1;
    },
    plotCountries() {
      // Not Visited (no hover)
      // Create map polygon series
      var polygonSeriesDisbaled = this.chart.series.push(
        new am4maps.MapPolygonSeries()
      );
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
      var polygonSeries = this.chart.series.push(
        new am4maps.MapPolygonSeries()
      );
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

      this.focusOnCountry(polygonSeries, "SG");
    },
    plotStatesUSA() {
      let usaSeries = this.chart.series.push(new am4maps.MapPolygonSeries());
      usaSeries.geodata = am4geodata_usaLow;

      // Display only needed states
      usaSeries.include = visited_states_usa;

      let usPolygonTemplate = usaSeries.mapPolygons.template;
      usPolygonTemplate.tooltipText = "{name}";
      // usPolygonTemplate.fill = this.chart.colors.getIndex(1);
      // usPolygonTemplate.nonScalingStroke = true;
      usPolygonTemplate.fill = am4core.color("#74B266"); // old color: #74B266
      usPolygonTemplate.stroke = am4core.color("#000000"); // outline color
      usPolygonTemplate.strokeWidth = 0.5;
      usPolygonTemplate.strokeOpacity = 0.8;

      // Hover state
      let hs = usPolygonTemplate.states.create("hover");
      hs.properties.fill = am4core.color("#367B25");
    },
    plotStatesIndia() {
      let indiaSeries = this.chart.series.push(new am4maps.MapPolygonSeries());
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
    },
    focusOnCountry(polygonSeries, countryCode) {
      // TODO: error or return on else? or leave behaviour as is?
      if (this.latitude_map.has(countryCode)) {
        this.chart.deltaLongitude = this.latitude_map.get(countryCode);
      }
      this.chart.events.on("ready", () => {
        // Zoom
        // NOTE: this will only work when the polygon is included in the series (USA and INDIA -> removed from global series)
        // this.chart.events.on("ready", function() {
        var countryPolygon = polygonSeries.getPolygonById(countryCode);

        // Pre-zoom
        this.chart.zoomToMapObject(countryPolygon);

        // Set active state
        setTimeout(function() {
          countryPolygon.isActive = true;
        }, 1000);
      });
    },
    rotateTo(long, lat) {
      if (this.animation) {
        this.animation.stop();
      }
      this.animation = this.chart.animate(
        [
          {
            property: "deltaLongitude",
            to: long
          },
          {
            property: "deltaLatitude",
            to: lat
          }
        ],
        2000
      );
    },
    animate() {
      // this.chart.goHome(); // to reset zoom
      setTimeout(() => {
        this.animation = this.chart.animate(
          { property: "deltaLongitude", to: 100000 },
          20000000
        );
      }, 3000);

      this.chart.seriesContainer.events.on("down", () => {
        if (this.animation) {
          this.animation.stop();
        }
      });

      this.chart.seriesContainer.events.on("up", () => {
        this.animation = this.chart.animate(
          { property: "deltaLongitude", to: 100000 },
          20000000
        );
      });
    }
  },
  mounted: function() {
    // Create map instance
    this.chart = am4core.create(this.$refs.chartdiv, am4maps.MapChart);
    this.setupChart();
    this.plotBgAndLines();
    this.plotCountries();
    this.plotStatesUSA();
    this.plotStatesIndia();
    // this.animate();
  }
};
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
