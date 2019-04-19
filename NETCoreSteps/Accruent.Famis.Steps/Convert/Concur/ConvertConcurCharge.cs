using Famis.Model;
using StepCore;

namespace Accruent.Famis.Steps.Convert.Concur {
    [StepDescription("convert_concur_charge")]
    public class ConvertConcurCharge : Step{
        
        [Input(Description = "The pipe delimited charge from the concur file")] public string ChargeLine { get; set; }
        [Output(Description = "A FAMIS Other cost converted from a concur charge")] public OtherCost Record { get; set; }
        
        public override void Execute() {
            var components = ChargeLine.Split("|");
            var otherCost = new OtherCost();
            otherCost.InvoiceNumber = components[1];
            Record = otherCost;
        }
    }
}