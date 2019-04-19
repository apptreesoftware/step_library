using System.Collections.Generic;
using System.Threading.Tasks;
using Famis;
using StepCore;

namespace Accruent.Famis.Steps.Fetch {
    [StepDescription("fetch_records", Description = "Fetches a list of records from famis")]
    public class FetchRecords : ServiceStep {
        [Input(Description = "the endpoint to use for this record")]
        public string Endpoint { get; set; }
        
        [Input(Description = "optional filter to use when fetching records")]
        public string Filter { get; set; }
        
        [Input(Description = "optional expand value")]
        public string Expand { get; set; }
        
        [Input(Description = "optional page offset, default is 0")]
        public int Offset { get; set; }

        [Input(Description = "optional maximum number of records to return")]
        public int Limit { get; set; }
        
        [Output(Description = "the fetched records")]
        public List<object> Records { get; set; }
        
        public override async Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            if (Limit == 0) {
                Limit = -1;
            }

            Records = await service.GetRecords(Endpoint, Filter, Expand, Offset, Limit);
        }
    }
}
