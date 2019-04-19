using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {

    public class AccountSegmentNPFA {
        public int Id { get; set; }
        public string Name { get; set; }
        public bool Active { get; set; }
        public bool IncludeInValidationFlag { get; set; }
        public int TabOrder { get; set; }
        public int CharacterLimit { get; set; }
        public int CaseRestriction { get; set; }
        public string ExternalId { get; set; }
    }

}
