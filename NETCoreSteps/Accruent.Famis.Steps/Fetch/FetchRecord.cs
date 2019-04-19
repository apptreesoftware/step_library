using System;
using System.Threading.Tasks;
using Famis;
using StepCore;

namespace Accruent.Famis.Steps.Fetch {
    [StepDescription(
        "fetch_record", Description = "Fetches a single record from famis")]
    public class FetchRecord : ServiceStep {
        [Input(Description = "the endpoint to use for this record")]
        public string Endpoint { get; set; }

        [Input(Description = "the filter used to fetch record")]
        public string Filter { get; set; }

        [Input(Description = "optional expand value")]
        public string Expand { get; set; }

        [Output(Description = "the fetched record")]
        public object Record { get; set; }

        public override async Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            var record = await service.GetRecord(Endpoint, Filter, Expand);
            Record = record;
        }
    }
}