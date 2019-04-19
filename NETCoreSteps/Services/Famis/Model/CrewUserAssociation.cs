using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {

    public class CrewUserAssociation {
        public int Id { get; set; }
        public int UserId { get; set; }
        public string UserExternalId { get; set; }
        public int CrewId { get; set; }
        public string CrewExternalId { get; set; }
        public DateTime UpdateDate { get; set; }
        public int UpdatedById { get; set; }
        public string UpdatedByExternalId { get; set; }
        public double? Rate { get; set; }
        public double? OT { get; set; }
        public double? DT { get; set; }
        public bool UseCrewRatesFlag { get; set; }
        public bool DefaultCrewFlag { get; set; }
        public bool CrewLeaderFlag { get; set; }
    }

}
