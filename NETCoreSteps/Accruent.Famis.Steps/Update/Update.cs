using System.Threading.Tasks;
using Famis;
using StepCore;

namespace Accruent.Famis.Steps.Update
{
    [StepDescription("update")]
    public class Update : FamisUpsert
    {
        public override async Task ExecuteAsync() {
            var service = new Service(Url, Username, Password);
            var resp = await service.UpdateRecord(Endpoint, Object, Id);
            Message = resp.Message;
            Record = resp.Object;
            Success = resp.Success;
        }
    }
}