using System.Data;
using System.Threading.Tasks;
using Famis;
using StepCore;

namespace Accruent.Famis.Steps.Create
{
    [StepDescription("create")]
    public class Create : FamisUpsert
    {
        public override async Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            var resp = await service.CreateRecord(Endpoint, Object);
            Message = resp.Message;
            Record = resp.Object;
            Success = resp.Success;
        }
    }
}